package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const GOROUTINES=1000000

func doIt(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for i:=0;i<GOROUTINES;i++ {
		wg.Add(1)
		go doIt(&wg)
	}
	spinUptime := time.Since(start)
	fmt.Printf("Spun up %d goroutines in %v. %v per goroutine\n", GOROUTINES,spinUptime, spinUptime/GOROUTINES )
	memstat := PrintMemUsage()
	fmt.Printf("Memory usage is %d bytes per goroutine\n", memstat.Sys/GOROUTINES)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %v\n", elapsed)
}




func PrintMemUsage() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	return m
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}