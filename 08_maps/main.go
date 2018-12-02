package main

import "fmt"

func main() {
	//	// Define map
	//	emails := make(map[int]string)
	//	// Assign kv
	//	emails[1] = "alice@gmail.com"
	//	emails[2] = "bob@gmail.com"
	//	emails[3] = "charlie@gmail.com"

	// Declare map and add kv
	emails := map[int]string{1: "alice@gmail.com", 2: "bob@gmail.com", 3: "charlie@gmail.com"}
	emails[4] = "darek@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[2])

	// delete key
	delete(emails, 2)
	fmt.Println(emails)
}
