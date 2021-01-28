package main

import (
	"context"
	"fmt"
	"time"
)

func printName(ctx context.Context, name string) {
	minTime := 500 * time.Millisecond
	if deadlineTime, ok := ctx.Deadline(); ok {
		if time.Until(deadlineTime) < minTime {
			fmt.Println("Time is too short to complete task")
			return
		}
	}
	select {
	case <-time.After(minTime):
		fmt.Println("Name received is:", name)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	printName(ctx, "John")

	fmt.Println("Main exited.")
}
