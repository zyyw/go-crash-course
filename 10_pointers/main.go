package main

import "fmt"

func main() {

}

func PointerExample1() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}

func PointerExample2() {
	var a int = 2
	// 注意这里是 *int, 不是 int* (c 语言是 int* )
	var pa *int = &a
	*pa = 3
	fmt.Println(a) // output: 3
}

// swap 的 三种形式
// 1. the arguments passed is passed through value, won't be swapped
func swap(a, b int) {
	a, b = b, a
}

// 2. use pointer; arguments passed in will be swapped
func swap2(a, b *int) {
	*a, *b = *b, *a
}

// 3. use multiple return values
func swap3(a, b int) (int, int) {
	return b, a
}
