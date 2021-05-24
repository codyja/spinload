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

//func loadCpu(c float64) {
//	cpuLimiter := &cpulimit.Limiter{
//		//MaxCPUUsage: float64(l.CpuPercent),
//		MaxCPUUsage: c,
//		SwitchPeriod: 100 * time.Millisecond,
//		MeasurePeriod: 25 * time.Millisecond,
//		MeasureDuration: 200 * time.Millisecond,
//
//	}
//
//	cpuLimiter.Start()
//	defer cpuLimiter.Stop()
//
//	done := make(chan int)
//
//	for i := 0; i < runtime.NumCPU(); i++ {
//		go func() {
//			for {
//				select {
//				case <-done:
//					return
//				case <-cpuLimiter.H:
//					// over desired cpu usg, pause for a bit to cool down
//					time.Sleep(time.Millisecond * 50)
//				default:
//					// infinite loop
//				}
//			}
//		}()
//	}
//
//	// keep socket open for 2000ms to simulate a busy web server, then close channel
//	time.Sleep(time.Second * 3)
//	close(done)
//}

//func (l *load) loadGen(w http.ResponseWriter, r *http.Request) {
func loadGen(w http.ResponseWriter, r *http.Request) {

	//query := r.URL.Query()

	cpuLimiter := &cpulimit.Limiter{
		//MaxCPUUsage: float64(10),
		//MaxCPUUsage: float64(l.CpuPercent),
		//MaxCPUUsage: float64(50),
		SwitchPeriod: 100 * time.Millisecond,
		MeasurePeriod: 25 * time.Millisecond,
		MeasureDuration: 200 * time.Millisecond,

	}

	//cpuPercent := query.Get("cpu")
	//if cpuPercent != "" {
	//	cpuLimiter.MaxCPUUsage = parseFloat(cpuPercent, 64)
	//
	//	// start up cpu load
	//	//loadCpu(parseFloat(cpuPercent, 64))
	//}

	//fmt.Printf("Limiter = %f", cpuLimiter.MaxCPUUsage)


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
					time.Sleep(time.Millisecond * 1000)
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
	fmt.Fprintf(w, "Your load has been delivered... \n")
}

func main() {
	//cpuPercent := flag.Int("cpu-percent", 50, "Percent limit of CPU to load. Default: 50")
	//flag.Parse()

	//loadHandler := &load{CpuPercent: *cpuPercent}

	http.HandleFunc("/load", loadGen)

	log.Fatal(http.ListenAndServe(":1986", nil))
}
