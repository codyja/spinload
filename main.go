package main

import (
	"fmt"
	"github.com/codyja/go-cpulimit"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"
)

//type load struct {
//	CpuPercent int
//}

func parseFloat (s string, bitSize int) (float64) {
	f, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func loadGen(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	cpuLimiter := &cpulimit.Limiter{
		SwitchPeriod: 100 * time.Millisecond,
		MeasurePeriod: 25 * time.Millisecond,
		MeasureDuration: 200 * time.Millisecond,

	}

	cpuPercent := query.Get("cpu")

	var brakeTime time.Duration

	if cpuPercent != "" {
		percent := parseFloat(cpuPercent, 64)
		cpuLimiter.MaxCPUUsage = percent

		switch {
		case percent <= 20:
			brakeTime = 550 * time.Millisecond
		case percent <= 40:
			brakeTime = 250 * time.Millisecond
		case percent <= 60:
			brakeTime = 200 * time.Millisecond
		case percent <= 80:
			brakeTime = 5 * time.Millisecond
		case percent <= 100:
			brakeTime = 10 * time.Millisecond
		}

	}

	cpuLimiter.Start()
	defer cpuLimiter.Stop()

	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				case <-cpuLimiter.H:
					// over desired cpu usg, pause for a bit to cool down
					//time.Sleep(time.Millisecond * 550)
					time.Sleep(brakeTime)
				default:
				}
			}
		}()
	}

	// keep socket open for 2000ms to simulate a busy web server, then close channel
	time.Sleep(time.Second * 3)
	close(done)

	// write msg to http.ResponseWriter and close
	fmt.Fprintf(w, "Your load has been delivered... \n")
}

func main() {
	//cpuPercent := flag.Int("cpu-percent", 50, "Percent limit of CPU to load. Default: 50")
	//flag.Parse()

	//loadHandler := &load{CpuPercent: *cpuPercent}

	http.HandleFunc("/load", loadGen)

	log.Fatal(http.ListenAndServe(":1986", nil))
}
