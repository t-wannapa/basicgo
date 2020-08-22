package main

import "fmt"

type Day int

func (d Day) Today() string {
	return fmt.Sprintf("today : %d", d)
}

func main() {
	d := Day(2)
	today := d.Today()
	fmt.Println(today)
}
