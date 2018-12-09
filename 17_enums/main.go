package main

import (
	"fmt"
)

const (
	// UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	// TRACE logs everything
	TRACE // 1
	// INFO logs Info, Warnings and Errors
	INFO // 2
	// WARNING logs Warning and Errors
	WARNING // 3
	// ERROR just logs Errors
	ERROR // 4
)

// Level holds the log level.
type Level int

// SetLogLevel ...
func SetLogLevel(level Level) {
	switch level {
	case TRACE:
		fmt.Println("trace")
		return

	case INFO:
		fmt.Println("info")
		return

	case WARNING:
		fmt.Println("warning")
		return
	case ERROR:
		fmt.Println("error")
		return

	default:
		fmt.Println("default")
		return

	}
}

// Suite ...
type Suite int

const (
	// Spades ...
	Spades Suite = iota
	// Hearts ...
	Hearts
	// Diamonds ...
	Diamonds
	// Clubs ...
	Clubs
)

func (s Suite) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs"}[s]
}

func main() {
	// SetLogLevel(INFO)
	// fmt.Println(INFO)

	fmt.Println(Hearts)
}
