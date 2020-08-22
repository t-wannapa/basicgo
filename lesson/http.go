package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "anuchit"}`)
	w.Header().Add("content-type", "application/json")
	w.Write(resp)
}
func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("starting...")
	err := http.ListenAndServe(":1234567890", nil)
	log.Println(err)
	fmt.Println("end")
}
