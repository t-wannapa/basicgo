package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println("address i ", &i)

	p := &i // point to i
	fmt.Println("p value inside address ", *p)
	*p = 21                    // set i through the pointer
	fmt.Println("i value ", i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
