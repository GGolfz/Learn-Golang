package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("%#v\n", langs[0:2])
	fmt.Printf("%#v\n", langs[1:3])
	fmt.Printf("%#v\n", langs[0:])
	fmt.Printf("%#v\n", langs[:2])
	printSlices(langs)
	cords := [4]string{"Am", "B", "G", "F#"}
	printSlices(cords[:])
}

func printSlices(ns []string) {
	fmt.Printf("Print Slice: ns: %#v\n", ns)
}
