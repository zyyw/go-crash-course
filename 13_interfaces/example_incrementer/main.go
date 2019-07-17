package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

// dosen't have to a struct type to implement an interface
// basically, any type of user defined can implement an interface
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
