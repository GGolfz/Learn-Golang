package main

import (
	"fmt"

	"github.com/ggolfz/igapp/time"
	"github.com/ggolfz/igapp/user"
)

func main() {
	user.Info("GGolfz", "Hello", 20)
	fmt.Println(time.Today())
}
