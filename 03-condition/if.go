package main

import "fmt"

func main() {
	num := 14
	if num%2 == 0 {
		fmt.Println("Even")
	} else if num == 15 {
		fmt.Println("Special")
	} else {
		fmt.Println("Odd")
	}
}
