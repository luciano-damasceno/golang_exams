package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	// ctx := context.Background()
	// _ = ctx

	const numJobs = 10000

	jobsCh := make(chan int, numJobs)
	jobsCompletedCh := make(chan int, numJobs)

	for i := 0; i < 3; i++ {
		go loopMemTest(jobsCh, jobsCompletedCh)
	}

	for j := 1; j <= numJobs; j++ {
		jobsCh <- j
	}
	close(jobsCh)

	for a := 1; a <= numJobs; a++ {
		<-jobsCompletedCh
	}

	funcName := "Main"
	PrintMemUsage(&funcName)

	fmt.Printf("Num of CPUs: %v\n", runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Processes took %v\n", elapsedTime)
}

func loopMemTest(jobsCh <-chan int, jobsCompletedCh chan<- int) {
	for j := range jobsCh {
		if j == len(jobsCh) {
			funcIndex := strconv.Itoa(j)
			PrintMemUsage(&funcIndex)
		}
		jobsCompletedCh <- j
	}
	// fmt.Printf("\tGoroutine: %d\n", index)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garbage collection cycles completed.
func PrintMemUsage(index *string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	if index != nil {
		fmt.Printf("Function index: %v\tAlloc = %v KiB\tTotalAlloc = %v KiB\tSys = %v KiB\tNumGC = %v\tNum Goroutines: %v\n", *index, bToKb(m.Alloc), bToKb(m.TotalAlloc), bToKb(m.Sys), m.NumGC, runtime.NumGoroutine())
	}
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
