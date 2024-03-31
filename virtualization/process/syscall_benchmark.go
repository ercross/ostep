package main

import (
	"fmt"
	"syscall"
	"time"
)

// syscallBenchmark measures time-cost of a system call
//
// Note that result may vary for different system resources
// (e.g cpu clock-cycle rating and memory)
func syscallBenchmark() {
	iterations := 10000
	start := time.Now()

	// do a 0-byte read iterations time
	for i := 0; i < iterations; i++ {
		_, err := syscall.Read(0, []byte{})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	end := time.Now()

	// Calculate the time taken per iteration
	totalTime := end.Sub(start)
	avgTime := totalTime / time.Duration(iterations)

	fmt.Println("Average time per 0-byte read operation:", avgTime)
}
