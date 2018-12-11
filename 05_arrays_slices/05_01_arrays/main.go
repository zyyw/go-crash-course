package main

import "fmt"

/*
Note in-place:
1. A slice has three components: a pointer, a length, and a capacity.

2. This is an array initialization / literal:
a := [...]int{0, 1, 2, 3}
This is a slice initialization / literal:
s := []int{0, 1, 2, 3}

3.arrays are comparable, but slice are not. So we cannot use == to test
whether two slices contain the same elements.
For two slice of bytes, to compare, use bytes.Equal();
for other types of slices, implement it yourself.
*/
func main() {
	runArrays()
}

func runArrays() {
	// Arrays
	var fruitArr [2]string
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	// Declare and assign
	fruitArr2 := [2]string{"Apple", "Banana"}
	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[1])
}
