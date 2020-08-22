package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

func main() {
	v := Vertext{1, 2}

	p := &v
	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", *p)
	fmt.Printf("%#v\n", (*p).X)
}
