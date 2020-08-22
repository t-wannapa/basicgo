package main

import "fmt"

type Samsung struct {
	Version string
}

func (s Samsung) Info() string {
	return fmt.Sprintf("info: %s", s.Version)
}

type OnePlus struct {
	Samsung
	Version string
}

func main() {
	s := Samsung{Version: "s10+"}

	fmt.Println(s.Info())

	o := OnePlus{}
	o.Version = "7 Pro"

	fmt.Println(o.Info())
}
