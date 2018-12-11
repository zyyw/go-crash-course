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
	// runSlices2()
	// runAppend()

	// reverse a slice of ints
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(nums)
	// reverse(nums)
	// fmt.Println(nums)

	// rotate a slice of int to the left by 2 elements
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	rotate(nums)
	fmt.Println(nums)
}

func runSlices() {
	// Slice
	fruitSlice := []string{"Apple", "Banana", "Cherry", "Durian"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:])
	fmt.Println(fruitSlice[:1])
	fmt.Println(fruitSlice[1:3])
	fmt.Println(len(fruitSlice))
}

// this is an example from "The Go Programming Language"
func runSlices2() {
	months := []string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println("len of summer:", len(summer))
	fmt.Println("cap of summer:", cap(summer))
	fmt.Println("len of months:", len(months))
	fmt.Println("cap of months:", cap(months))
}

func runAppend() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	fmt.Println(x)
	x = append(x, x...)
	fmt.Println(x)
}

// reverse a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate a slice to the left by n elements
func rotate(s []int) {
	// here n = 2
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}
