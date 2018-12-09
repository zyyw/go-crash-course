package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 世界")
}

func CreateWebServer() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}

// Fetch prints the content found at a URL
func FetchPrint() {
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
	fmt.Println(resp.Status)
}

func FetchCopyPrint() {
	if len(os.Args) < 2 {
		fmt.Println("please provide with Url")
		os.Exit(1)
	}
	url := os.Args[1]
	if strings.HasPrefix(url, "http://") == false {
		url = "http://" + url
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}

// FetchWebConcurrently Concurrently fetch web
func FetchWebConcurrently(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resource
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}

// RunFetchConcurrency driver for FetchWebConcurrently()
func RunFetchConcurrency() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go FetchWebConcurrently(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	// CreateWebServer()
	// FetchPrint()
	// FetchCopyPrint() // go run main.go http://gopl.io
	RunFetchConcurrency() // go run main.go https://golang.org http://gopl.io https://godoc.org
}
