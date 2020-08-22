package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	printSlice(s)

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is", s)
	}
}

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
