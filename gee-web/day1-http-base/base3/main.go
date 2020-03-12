package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {

	s := gee.New()
	s.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	s.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Headers [%q] = %q\n", k, v)
		}
	})

	s.Run(":9995")
}
