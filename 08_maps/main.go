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

	// how to determine whether a key is in a map or not
	if _, ok := emails[2]; !ok {
		fmt.Println("The key 2 is NOT in the emails map.")
	}

	if _, ok := emails[4]; ok {
		fmt.Println("The key 4 is in the email map.")
	}
}

// To determine whether two maps contain the same keys and the same associated values
func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
