package main

import "fmt"

func main() {
	var v interface{}
	v = 1
	fmt.Printf("%T %#v\n", v, v)

	v = "1"
	fmt.Printf("%T %#v\n", v, v)

	v = map[string]int{
		"test": 1,
	}
	fmt.Printf("%T %#v\n", v, v)

	v = []string
	fmt.Prinf("%T %#v\n", v, v)
}
