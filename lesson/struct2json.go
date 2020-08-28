package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	t := Todo{
		ID:     1,
		Title:  "pay credit card",
		Status: "completed",
	}
	b, err := json.Marshal(t)
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Println(err)
}
