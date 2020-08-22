package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordSlice := strings.Split(s, " ")

	var wordMap map[string]int
	wordMap = make(map[string]int)
	for _, word := range wordSlice {
		wordMap[strings.Trim(word, ",.")] += 1
	}

	return wordMap
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}
