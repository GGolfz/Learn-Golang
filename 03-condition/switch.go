package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday
	switch today {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
		fallthrough // enable falling to below case
	default:
		fmt.Println("Today is others day")
	}
}
