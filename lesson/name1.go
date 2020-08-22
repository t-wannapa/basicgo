package main

import (
	"fmt"
	"math"
)

func compute1(fn func(float64, float64) float64) float64 {
	return fn(5, 12)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 2)
}
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(2, 2))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
