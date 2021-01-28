package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Currently No. of Goroutine running:", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for i := 1; i <= 100; i++ {
		go func(n int) {
			fmt.Println("Running Goroutine no. :", n)

			for {
				select {
				case <-ctx.Done():
					fmt.Println("Ending Goroutine no.:", n)
					return
				default:
					fmt.Println("Still running Goroutine no. :", n)
					time.Sleep(100 * time.Millisecond)
				}
			}

		}(i)
	}

	fmt.Println("Sleeing for 5seconds")
	time.Sleep(time.Millisecond)
	cancel()
	fmt.Println("Sleeing for 2seconds")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Currently No. of Goroutine running:", runtime.NumGoroutine())
	fmt.Println("End....")
}
