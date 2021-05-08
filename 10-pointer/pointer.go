package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("You are %s\n", me)
	fmt.Println(&me) // address of me variable
	fmt.Printf("%T\n", &me)
	// change value using pointer
	var addr *string = &me
	*addr = "penguin"
	fmt.Printf("You are %s\n", me)
}
