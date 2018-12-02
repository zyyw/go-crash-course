package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getMultiply(num1, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(greeting("Brad"))
	fmt.Println(getSum(3, 4))
	fmt.Println(getMultiply(3, 4))
}
