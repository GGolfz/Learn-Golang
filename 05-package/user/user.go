package user

import "fmt"

func Info(name string, msg string, age int) {
	fmt.Printf("My name is %v\n", name)
	fmt.Printf("My msg is %v\n", msg)
	fmt.Printf("I'm %d years old\n", age)
}
