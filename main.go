package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, NewMessage())
	log.Print("mylog")
	log.Print("mylog2")
}

func NewMessage() string {
	return "Hello, World"
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
