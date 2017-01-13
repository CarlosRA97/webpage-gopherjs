package main

import (

	"fmt"
	"net/http"

	//"github.com/gopherjs/gopherjs/js"
)

func main() {
	// Servidor web
	http.HandleFunc("/", func (response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello there")
	})

	http.ListenAndServe(":8080", nil)
}