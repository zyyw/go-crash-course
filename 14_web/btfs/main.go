package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("main server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "BTFS node: main server")
	})
	//http.HandleFunc("/", handler)
	http.ListenAndServe(":3001", nil)
}
