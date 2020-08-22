package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	printSlice(s)
	s = append(s, 99, 33, 44, 55)
	s = s[:2]
	printSlice(s)

	s = s[1:]

	printSlice(s)
}

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap%d\n", len(s), cap(s))
}
