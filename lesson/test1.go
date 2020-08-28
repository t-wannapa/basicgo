package main

import "fmt"

type Todo struct {
	ID int
}

func test(t *Todo) {
	fmt.Println(t.ID)
}

func main() {

	t := Todo{ID: 1}
	fmt.Println(&t.ID)
	test(&t)
	fmt.Printf("%#v\n", t)

}
