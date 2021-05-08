package main

import "fmt"

var plus = func(a int, b int) int { return a + b }

func main() {
	var name string = "Golfz"
	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)
	p := plus(1, 2)
	fmt.Println("1+2 = ", p)
	fmt.Printf("%T\n", plus)
	newplus := func(a int, b int) int { return a + b }
	fmt.Println("2+3 = ", newplus(2, 3))
	fmt.Println("result: ", cal(4, plus, 5))
	fmt.Println("result: ", cal(4, minus, 5))
}

// Call function inside function
func cal(a int, op func(int, int) int, b int) int {
	return op(a, b)
}

func minus(a, b int) int {
	return a - b
}
