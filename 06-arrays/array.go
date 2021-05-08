package main

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("show all: %#v\n", ns)
}
func main() {
	var langs = [4]string{"golang", "python", "javascript"}
	fmt.Printf("langs: %v\n", langs)
	fmt.Printf("langs: %#v\n", langs)
	fmt.Printf("value at index 1: %v\n", langs[1])
	langs[1] = "pythonistra"
	fmt.Printf("value at index 1: %v\n", langs[1])
	showAll(langs)
}
