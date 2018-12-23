package main

import (
	"fmt"
	"time"
)

func main() {
	captureIterationVariable()
	fmt.Println()
	noCaptureIterationVariable()
	fmt.Println()
}

// capturing iteration variables
// print: "12345"
func captureIterationVariable() {
	nums := []int{1, 2, 3, 4, 5}
	var printNum []func()
	for _, d := range nums {
		value := d
		printNum = append(printNum, func() {
			fmt.Print(value)
		})
	}
	time.Sleep(3 * time.Second)
	for _, val := range printNum {
		val()
	}
}

// not capturing iteration variables
// print: "55555"
func noCaptureIterationVariable() {
	nums := []int{1, 2, 3, 4, 5}
	var printNum []func()
	for _, d := range nums {
		printNum = append(printNum, func() {
			fmt.Print(d)
		})
	}
	time.Sleep(3 * time.Second)
	for _, val := range printNum {
		val()
	}
}
