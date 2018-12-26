package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(apply(pow, 3, 2))

	// this code below within main is the equvalent with the one above
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 2))
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

// normal function
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported operation: %s" + op)
	}
}

// implement eval with functional programming style
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}
