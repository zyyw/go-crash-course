package main

import "fmt"

func main() {
	//	// Arrays
	//	var fruitArr [2]string
	//	// Assign values
	//	fruitArr[0] = "Apple"
	//	fruitArr[1] = "Banana"

	//	// Declare and assign
	//	fruitArr := [2]string{"Apple", "Banana"}
	//	fmt.Println(fruitArr)
	//	fmt.Println(fruitArr[1])

	// Slice
	fruitSlice := []string{"Apple", "Banana", "Cherry", "Durian"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:])
	fmt.Println(fruitSlice[:1])
	fmt.Println(fruitSlice[1:3])
	fmt.Println(len(fruitSlice))
}
