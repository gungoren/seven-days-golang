package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header [%q] = %q\n", k, v)
		}
	})

	http.ListenAndServe(":9995", nil)
}
