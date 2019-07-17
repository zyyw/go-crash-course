package main

import (
	"fmt"
	"reflect"
)

func main() {
	JustCallMethodSet()
}

type S struct{}
type T struct {
	S
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}

func JustCallMethodSet() {
	var t T

	methodSet(t)
	println("---------------")
	methodSet(&t)
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
