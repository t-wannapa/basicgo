package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s, _ := i.(string) // assert
	fmt.Println(s)

	if s, ok := i.(string); ok {
		fmt.Println(s)
	}
}
