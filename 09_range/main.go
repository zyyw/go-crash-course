package main

import "fmt"

func main() {
	ids := []int{34, 56, 21, 48, 11, 73}

	// Loop through ds
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %+v\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	// Range with map
	emails := map[int]string{1: "alice@gmail.com", 2: "bob@gmail.com", 3: "charlie@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%d: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("ID:", k)
	}

	for _, v := range emails {
		fmt.Println("email:", v)
	}
}
