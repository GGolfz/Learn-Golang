package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func main() {
	u := User{Username: "golang"}
	fmt.Printf("%#v\n", u)
	u.FullName = "Go Language"
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.Username)
}
