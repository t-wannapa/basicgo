package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	fmt.Printf("%#v\n", primes)
	fmt.Printf("%#v\n", s)

	s[0] = 6
	fmt.Printf("%#v\n", primes)
	fmt.Printf("%#v\n", s)
}