package main

import "fmt"

func main() {
	var langs []string
	fmt.Printf("langs: %#v\n", langs)
	if langs == nil {
		fmt.Printf("This is empty slice\n")
	} else {
		fmt.Printf("This is not empty slice\n")
	}
}
