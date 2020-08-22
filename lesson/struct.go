package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	t := v
	t.X = 15
	t.Y = 20
	fmt.Println(v)
	fmt.Println(t)
}
