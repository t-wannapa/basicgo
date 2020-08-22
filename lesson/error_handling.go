package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./defer.gox")
	if err != nil {
		// handle error
		log.Fatal(err)
		//return err
	}
	defer f.Close()
	// continue
	fmt.Println("good bye")
}
