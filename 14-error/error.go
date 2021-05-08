package main

import (
	"errors"
	"fmt"
	"log"
)

func ReadFile(name string) (string, error) {
	// read file ....
	return "", errors.New("file not found")
}
func main() {
	var err1 error = errors.New("File not found")
	fmt.Printf("%#v\n", err1)

	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
