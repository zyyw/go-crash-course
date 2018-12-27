package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// part 1: Encode
	// s, just a string of Ascii characters
	s := "Hello"
	// H: 72 in Ascii table
	// e: 101 in Ascii table
	// l: 108 in Ascii table
	// o: 111 in Ascii table
	// convert a string to byte slice
	sBytes := []byte(s)
	fmt.Println("sBytes", sBytes) // output: [72 101 108 108 111]
	// encode the byte slice to hexString
	hexEncodedString := hex.EncodeToString(sBytes)
	// 72 is "48" in hex representation
	// 101 is "65" in hex representation
	// 108 is "6c" in hex representation
	// 111 is "6f" in hex representation
	fmt.Println("hexEncodedString", hexEncodedString) //output: "48656c6c6f"

	// encode a byte slice to a byte slice with hexadecimal encoding
	// src byte slice: [72 101 108 108 111]
	// encode src with dexadecimal encoding, get string representation of the encoded digest: "48656c6c6f"
	// the byte slice representation of the encoded digest is: [52 56 54 53 54 99 54 99 54 102]
	dst := make([]byte, hex.EncodedLen(len(sBytes)))
	hex.Encode(dst, sBytes)
	fmt.Println("dst", dst)         // output: dst [52 56 54 53 54 99 54 99 54 102]
	fmt.Println("str", string(dst)) // output: str 48656c6c6f

	// part 2: Decode
	// hexString is a string representation of hex
	const hexString = "48656c6c6f"
	// DecodeString, decode a hexString to a byte representation of the hexString
	decoded, _ := hex.DecodeString(hexString)
	fmt.Println("decoded", decoded) // output: [72 101 108 108 111]
	fmt.Printf("%s\n", decoded)     // output: "Hello"
	fmt.Println(string(decoded))    // output: "Hello"
}
