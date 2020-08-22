package main

import "fmt"

type Student struct {
	Name string
	ID   int
}

// func (s Student) String() string {
// 	return "Hello my Student"
// }

func (s Student) String() string {
	return fmt.Sprintf("Hellow %d : %s\n", s.ID, s.Name)
}

func main() {
	s := Student{ID: 35, Name: "Test_name"}
	fmt.Println(s)
}
