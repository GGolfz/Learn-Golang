package main

import "fmt"

// Global Scope
var name string = "GGolfz"

func main() {
	// Local Scope
	name2 := "GGolfz"
	fmt.Printf("name: %v\n", name)
	// print type
	fmt.Printf("type: %T\n", name)
	fmt.Printf("name2: %v\n", name2)
	var name3 string // Empty will be defult value
	// auto add quote
	fmt.Printf("name3: %q\n", name3)
	name = "new name"
	fmt.Printf("newname: %v", name)
}
