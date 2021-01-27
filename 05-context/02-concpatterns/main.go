package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for num := range c {
			out <- num
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	num := gen(1, 2, 3, 4, 5)
	out := sq(num)
	for {
		sqred, ok := <-out
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(sqred)
	}

	for n := range sq(sq(gen(2, 4))) {
		fmt.Println(n)
	}

	fmt.Println("Using fan out and fan in")

	in := gen(2, 3, 4, 5, 6, 7, 8)

	c1 := sq(in)
	c2 := sq(in)

	numChan := merge(c1, c2)
	for num := range numChan {
		fmt.Println(num)
	}

	fmt.Println("Exited")
}
