package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("%#v\n", status)
	fmt.Printf("%d\n", len(status))
	status[200] = "Okie"
	status[201] = "Created"
	fmt.Printf("%#v\n", status)
	value := status[200]
	fmt.Printf("%#v\n", value)
	// delete key
	delete(status, 201)
	fmt.Printf("%#v\n", status)
	fmt.Printf("%#v\n", status[201])
	if v, ok := status[201]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
	var m map[string]string
	if m == nil {
		fmt.Println("Nil Map")
	}
}
