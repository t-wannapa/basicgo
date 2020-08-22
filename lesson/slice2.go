package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[2])
	s[2] = 123
	fmt.Println(s)

	s = append(s, 4, 5, 6)
	fmt.Println(s)

	data := []int{1, 2, 3, 4}
	result := append(s, data...)
	fmt.Println(result)
	fmt.Printf("%d\n", cap(result))
}
