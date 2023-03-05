package main

import (
	"fmt"
	"net/http"
	"web-framework-gee-7days/gee"
)

func main() {
	engine := gee.NewEngine()
	engine.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	engine.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	engine.Run(":8080")
}
