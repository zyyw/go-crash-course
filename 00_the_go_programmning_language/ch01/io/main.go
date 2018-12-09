package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FindDupFromStdin() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// counts[input.Text()]++
		line := input.Text()
		if line == "N" {
			break
		}
		counts[line]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// here, counts is a map reference to the data structure created by make in main func
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func FindDupFromFile() {
	// find dup from file
	counts := make(map[string]int)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup from file: %v\n", err)
	}
	countLines(f, counts)
	f.Close()
	// output count results
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func FindDupFromFile2() {
	counts := make(map[string]int)
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup from file: %v\n", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func main() {
	// FindDupFromStdin()
	// FindDupFromFile()
	FindDupFromFile2()
}
