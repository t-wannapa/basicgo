package main

import "fmt"

type Person struct {
	name    string
	friends map[string]int
}

func (p Person) Walk() {
	fmt.Println(p.name, "walking")
}

func (p Person) Eat() {
	fmt.Println(p.name, "eating")
}

func (p Person) Geeting() {
	fmt.Println("hello", p.name)
}

func (p Person) Name() string {
	return p.name
}

// non pointer samatic value
// pointer
func (p Person) SetName(n string) {
	p.name = n
	p.friends = map[string]int{"test": 5}
	//p.friends["benz"] = 5
}

func main() {
	p := Person{name: "Tom", friends: map[string]int{
		"benz": 1,
	}}
	p.Walk()
	p.Eat()
	p.Geeting()

	//n := p.Name()
	fmt.Printf("%#v\n", p)

	p.SetName("KBTG")
	fmt.Println(p.Name())
	fmt.Printf("%#v\n", p)
}
