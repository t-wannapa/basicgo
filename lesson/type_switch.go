package main

import "fmt"

type Samsung struct {
	Version string
}

func main() {
	var i interface{} = Samsung{Version: "s10+"}
	v := i.(type)
	fmt.Printf("%T %#v\n", v, v)
	// switch v := i.(type) {
	// case int:
	// 	fmt.Printf("%T %d\n", v, v)
	// case string:
	// 	fmt.Printf("%T %s\n", v, v)
	// case Samsung:
	// 	fmt.Printf("%T %s Yeah, my new phone\n", v, v.Version)
	// default:
	// 	fmt.Println("Undefined type")
	// }
}
