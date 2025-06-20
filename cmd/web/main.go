package main

import (
	"fmt"
	"net/http"
	"github.com/pkini2002/go-course/pkg/handlers"
)

const portNumber = ":8081" 

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting server on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}