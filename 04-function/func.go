package main

import "fmt"

func main() {
	info()
	info2("Gopler")
	fmt.Printf(today())
	fmt.Print(swap("Left GGolfz ", " Right GGolfz "))
}

// without parameter
func info() {
	name := "GGolfz"
	fmt.Printf("My name is %v\n", name)
}

// with parameter
func info2(name string) {
	fmt.Printf("My name is %v\n", name)
}

// with return type
func today() string {
	return "Today"
}

// with more than one return value

func swap(x, y string) (string, string) {
	return y, x
}
