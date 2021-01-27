package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("No. of Goroutines running: ", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i <= 100; i++ {
		go func(n int) {
			fmt.Println("Running Goroutine: ", n)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("Working", n)
					time.Sleep(100 * time.Millisecond)
				}
				fmt.Println("GOROUTINE RUNNING CURRENTLY: ", runtime.NumGoroutine())
			}
		}(i)
	}

	fmt.Println("Sleeping for 5 seconds")
	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("Sleeping for 5 more seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("End No. of goroutine running now:", runtime.NumGoroutine())
}
