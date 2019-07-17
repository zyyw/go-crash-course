package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const TotalAccountNumber = 10

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Please provide the output filename of the generated keys, and the output file only")
	}
	filename := os.Args[1]
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	// defer to close file handler
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Failed to close f using f.Close(): %v", err)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < TotalAccountNumber; i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Post("https://api.trongrid.io/wallet/generateaddress", "", nil)
			if err != nil {
				log.Printf("Failed to curl POST: %v", err)
				return
			}
			// defer to close resp.Body
			defer func() {
				if err := resp.Body.Close(); err != nil {
					log.Printf("Failed to close http response body using resp.Body.Close(): %v", err)
				}
			}()
			body, err := ioutil.ReadAll(resp.Body)

			if _, err = f.WriteString(string(body)); err != nil {
				log.Printf("Failed to f.WriteString(): %v", err)
			}
		}()
	}
	wg.Wait()
}
