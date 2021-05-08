package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// while loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	// infinity loop
	// k := 0
	// for {
	// 	fmt.Println(k)
	// 	k++
	// }
	// loop through slice
	langs := []string{"golang", "python", "javascript", "r"}
	for l := 0; l < len(langs); l++ {
		value := langs[l]
		fmt.Println(l, ":", value)
	}

	for index, value := range langs {
		fmt.Println(index, " : ", value)
	}
	// loop without index
	for _, value := range langs {
		fmt.Println(value)
	}
	for index := range langs {
		fmt.Println(index)
	}
}
