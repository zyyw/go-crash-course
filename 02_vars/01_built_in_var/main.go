package main

import "fmt"

func main() {
	// vars types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var to declare variable
	// var name string = "Brad" // string can be omitted, becase the data type can be inferred from the right-hand side
	// var name = "Brad"
	name := "Brad"
	var age int32 = 37
	const isMale = true

	// short hand
	// username := "Brad"
	// email := "Brad@gmail.com"
	username, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isMale, username, email)
	fmt.Printf("%T\n", isMale)
	fmt.Printf("%t\n", isMale)
}
