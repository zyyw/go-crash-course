package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	const s = "48"
	decoded, _ := hex.DecodeString(s)
	fmt.Println(decoded)
	fmt.Printf("%s\n", decoded)
	fmt.Println(string(decoded))
}
