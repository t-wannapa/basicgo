package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["nong"] = "Anuchit0"
	fmt.Println(m["nong"])
}
