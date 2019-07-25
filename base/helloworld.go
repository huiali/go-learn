package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("Hello world")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "<h1>Hello world!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
