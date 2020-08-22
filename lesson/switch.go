package main

import "fmt"

func main() {
	os := "linux"

	switch os := Runtime.GOOS; os {
	case "linux":
		fmt.Println("it's linux")
		fallthrough
	case "windows":
		fmt.Println("it's windows")
	case "drawin":
		fmt.Println("it's drawin")
	default:
		fmt.Println("my os %s\n", os)
	}
}
