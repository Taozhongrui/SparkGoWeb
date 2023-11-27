package main

import (
	"example.com/m/v2/sparkgoweb"
	"fmt"
	"net/http"
)

func main() {
	router := sparkgoweb.Default()
	router.GET("/get", indexHandler)
	router.POST("/post", helloHandler)
	router.Run(":1234")
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
