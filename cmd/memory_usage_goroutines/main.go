package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	// ctx := context.Background()
	// _ = ctx

	const numJobs = 10000

	wg := sync.WaitGroup{}
	for i := 0; i <= numJobs; i++ {
		wg.Add(1)
		if i == numJobs {
			funcIndex := strconv.Itoa(i)
			PrintMemUsage(&funcIndex)
		}
		go loopMemTest(i, &wg)
	}
	wg.Wait()

	funcName := "Main"
	PrintMemUsage(&funcName)

	fmt.Printf("Num of CPUs: %v\n", runtime.NumCPU())

	elapsedTime := time.Since(startTime)
	fmt.Printf("Processes took %v\n", elapsedTime)
}

func loopMemTest(index int, wg *sync.WaitGroup) {
	defer wg.Done()
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
