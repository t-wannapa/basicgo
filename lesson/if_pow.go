package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("else v:", v)
	}

	return lim
}

func main() {
	fmt.Println(pow(5, 2, 26))

}
