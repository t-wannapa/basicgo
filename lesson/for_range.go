package main

import "fmt"

func main() {
	names := []string{"A", "B", "C", "D"}

	for _, name := range names {
		fmt.Print(name)
	}
}
