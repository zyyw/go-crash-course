package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://golang.org"
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

func WaitForServer(url string) error {
	const timeout = 10 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
