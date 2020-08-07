package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var sum int = 0
	var a int = 0
	var b int = 1
	return func() int {
		if sum == 0 {
			sum = 1
			return 0
		}
		sum = a + b
		a = b
		b = sum
		return sum
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
