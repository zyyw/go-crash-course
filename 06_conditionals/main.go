package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//	IfCompact()
	fmt.Println(SwitchCase2(56))
}

func IfElse() {
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
}

func IfCompact() {
	const filename = "sample.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// when contents is only declared in the if..else block, then it cannot be accessed out of that block
	// fmt.Printf("%s\n", contents)

	var contents2 []byte
	if contents2, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents2)
	}
	fmt.Printf("%s\n", contents2)
}

func SwitchCase1() {
	x := 5
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

// SwitchCase2 switch without variable, but having conditions in case statement
func SwitchCase2(score int) string {
	grade := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		grade = "F"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	case score <= 100:
		grade = "A"
	}
	return grade
}
