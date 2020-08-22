package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
