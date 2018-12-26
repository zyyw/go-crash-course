package main

import (
	"flag"
	"fmt"
)

func main() {
	FlagCollection()
}

func FlagCollection() {
	wordPtr := flag.String("word", "default value for word", "usage message for word")
	numPtr := flag.Int("num", 23, "usage message for num")
	boolPtr := flag.Bool("fork", false, "usage message for fork")

	var svar string
	flag.StringVar(&svar, "svar", "default value for svar", "usage message for svar")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

func FlagString() {
	username := flag.String("U", "", "Your username")
	password := flag.String("P", "", "Your password")

	flag.Parse()

	result := login(*username, *password)
	if result {
		fmt.Println("Login success")
	} else {
		fmt.Println("Login failed")
	}
}

func login(username, password string) bool {
	if username == "Mury" && password == "123456" {
		return true
	}
	return false
}
