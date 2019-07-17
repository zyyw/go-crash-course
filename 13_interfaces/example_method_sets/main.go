package main

import "fmt"

// implementing with values v.s. pointers
// * method set of value is all methods with value receivers
// * method set of pointer is all methods, regardless of receiver type

func main() {
	// this line below will fail, because Write() has a pointer receiver
	// var wc WriterCloser = myWriterCloser{}

	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

// define interfaces:
// Writer
// Closer
// WriterCloser
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

// define struct implements WriterCloser by implementing Write() and Close() respectively
type myWriterCloser struct{}

func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
