package main

import "fmt"

type Student struct {
	Name, Class string
}

var s = map[int]Student{
	1: Student{
		"xxx", "A",
	},
	2: Student{
		"zzz", "B",
	},
}

func main() {
	fmt.Println(s)
}
