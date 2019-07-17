package main

import "fmt"

// empty interface and type switches:
// the empty interface is nothing magic; it's just an inteface defined on the fly that has no methods on it
// every type in Go implements empty interface

func main() {

}

func simpleExample() {
	var v interface{} = 2
	switch v.(type) {
	case int:
		fmt.Println("v is an integer")
	case string:
		fmt.Println("v is a string")
	default:
		fmt.Println("unknown data type of v")
	}
}

// implement queue
// version 1: queue of type int only
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// version 2: queue of any type using empty interface
type Queue2 []interface{}

func (q *Queue2) Push(v interface{}) {
	*q = append(*q, v)
}
func (q *Queue2) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue2) IsEmpty() bool {
	return len(*q) == 0
}
