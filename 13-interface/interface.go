package main

import "fmt"

type Phone interface {
	Call(number string)
}
type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println(s.Name, "calling", number)
}
func (s Samsung) Answer() {
	fmt.Println(s.Name, "Hello There!")
}
func Dial(p Phone) {
	// Not found Answer because Answer not in interface phone
	p.Call("000-000-0000")
}
func main() {
	s := Samsung{
		Name: "S10",
	}
	Dial(s)
	s.Answer()
}
