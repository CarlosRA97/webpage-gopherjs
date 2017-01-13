package main

import (

	"fmt"
	"net/http"

	//"github.com/gopherjs/gopherjs/js"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	// Servidor web
	http.HandleFunc("/", func (response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello there")
	})

	if port == "" {
		port = "8080"
	}


	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}