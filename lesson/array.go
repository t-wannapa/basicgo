package main

import "fmt"

func show(s [3]string) {
	fmt.Println(s)
}

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1], a[2])

	primes := [6]int{2, 3, 5, 7, 13}
	fmt.Println(primes)

	show(a)
	show(primes)

}
