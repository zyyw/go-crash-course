package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.9
	var num = int(a)
	fmt.Println(num)
	fmt.Println(math.Ceil(a))
	fmt.Println(math.Floor(a))

	fmt.Println(math.Floor(a + 0.5))
}
