package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		select {
		case out <- n:
		case <-done:
			return out
		}
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			select {
			case out <- num * num:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for num := range c {
			select {
			case out <- num:
			case <-done:
				return
			}
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
	done := make(chan struct{})
	defer close(done)

	num := gen(done, 1, 2, 3, 4, 5)
	out := sq(done, num)

	for {
		sqred, ok := <-out
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(sqred)
	}

	for n := range sq(done, sq(done, gen(done, 2, 4))) {
		fmt.Println(n)
	}

	fmt.Println("Using fan out and fan in")

	in := gen(done, 2, 3, 4, 5, 6, 7, 8)

	c1 := sq(done, in)
	c2 := sq(done, in)

	numChan := merge(done, c1, c2)
	for num := range numChan {
		fmt.Println(num)
	}

	// Incase you decided to accept only 1, 2 or few values from merge func then
	// you can do it and it wont hang any goroutine , since defer done will be
	// called at the end that will trigger all goroutine to be returned.
	// Ex:
	// numChan := merge(done, c1, c2)
	// fmt.Println(<-numChan)
	// only 1 value gets accepted in above case.
	// Belore other values even gets started or already in process, when done gets
	// called it all goroutine will exit.
	// if u dont use done pattern then goroutine will not be exited and it will take
	// some resourses and also goroutines are not garbaged collected.

	fmt.Println("Exited")
}
