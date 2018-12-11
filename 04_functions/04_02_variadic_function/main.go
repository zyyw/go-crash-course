package main

import (
	"fmt"
)

// this is a variadic function taking an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print("The sum of ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(" is", total)
}

func main() {
	// variadic function can be called in the usual way with individual arguments
	sum(1, 2)    // "the sum of [1 2] is 3"
	sum(1, 2, 3) // "the sum of [1 2 3] is 6"

	// if you already have multiple args in a slice, apply them to a
	// variadic function using func(slice...) like this
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
