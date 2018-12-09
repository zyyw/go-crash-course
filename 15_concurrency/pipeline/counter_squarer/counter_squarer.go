package main

import (
	"fmt"
	"time"
)

/*
"The Go programming language", page 246
"Go Recipes", page 104
*/
func main() {
	Pipeline1()
	// Pipeline2()
	// RunDirectionalChannel()
	fmt.Println("YYY")
}

func Pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer
	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}

func Pipeline2() {
	naturals := make(chan int)
	squarers := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squarers <- x * x
		}
		close(squarers)
	}()

	// Printer
	for x := range squarers {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

// undirectional channel types
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
		time.Sleep(100 * time.Millisecond)
	}
}

func RunDirectionalChannel() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}
