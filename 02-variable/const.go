package main

import "fmt"

func main() {
	const day string = "Monday"
	fmt.Println("day:", day)
	// Const is constant cannot edit value

	const (
		Sunday = iota
		Monday
	) // auto create group of const variable and assign value by index
	fmt.Println("Sunday: ", Sunday)
	fmt.Println("Money: ", Monday)
}
