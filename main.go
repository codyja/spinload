package main

import (
	"flag"
	"fmt"
	"github.com/codyja/go-cpulimit"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

type StressHandler struct {
	CpuPercent int
}

func (sh *StressHandler) stressGen(w http.ResponseWriter, r *http.Request) {
	limiter := &cpulimit.Limiter{
		MaxCPUUsage: float64(sh.CpuPercent),
		SwitchPeriod: 100 * time.Millisecond,
		MeasurePeriod: 25 * time.Millisecond,
		MeasureDuration: 200 * time.Millisecond,

	}
	limiter.Start()
	defer limiter.Stop()

	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {

			for {
				select {
				case <-done:
					return
				case <-limiter.H:
					// over desired cpu usg, pause for a bit to cool down
					time.Sleep(time.Millisecond * 50)
				default:
					// infinite loop
				}

			}
		}()
	}

	// keep socket open for 2000ms to simulate a busy web server, then close channel
	time.Sleep(time.Second * 3)
	close(done)

	// write msg to http.ResponseWriter and close
	fmt.Fprintf(w, "Your stress has been delivered... \n")
}

func main() {
	cpuPercent := flag.Int("cpu-percent", 50, "Percent limit of CPU to stress. Default: 50")
	flag.Parse()

	stressHandler := &StressHandler{CpuPercent: *cpuPercent}

	http.HandleFunc("/stress", stressHandler.stressGen)

	log.Fatal(http.ListenAndServe(":1986", nil))
}
