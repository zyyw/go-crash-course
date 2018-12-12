package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide with URL")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", bytes)
}
