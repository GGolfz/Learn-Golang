package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!!!")
	// Format String
	fmt.Printf("Hello %s!!!\n", "GGolfz")
	// Format Integer
	fmt.Printf("Hello %d!!!\n", 50)
	// Format Any type
	fmt.Printf("Hello %v", "Test")
	fmt.Printf("Hello %v", 55)
	// Format more than one var
	fmt.Printf("Hello %v : %v!!!\n", "Ggolfz", 20)

}
