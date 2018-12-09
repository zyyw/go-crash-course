package main

import (
	"fmt"
)

func main() {
	// unbufferedChannel1()
	unbufferedChannel2()
}

// this will be dead lock
func unbufferedChannel1() {
	counter := make(chan int)
	counter <- 10
	fmt.Println(<-counter)
}

// this will not be dead lock
// because it exucutes the send operation by launching a new goroutine
// and the receive operation is executed in the main goroutine
func unbufferedChannel2() {
	counter := make(chan int)
	go func() {
		counter <- 10
	}()
	fmt.Println(<-counter)
}
