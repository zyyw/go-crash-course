package main

import "fmt"

func main() {
	//  if ... else
	x := 5
	y := 10
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is NOT less than %d\n", x, y)
	}

	// if ... else if ... else
	color := "green"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	// switch
	switch x {
	case 5:
		fmt.Printf("x is %d\n", 5)
	case 10:
		fmt.Printf("x is %d\n", 10)
	default:
		fmt.Println("x is NOT 5 or 10")
	}
}
