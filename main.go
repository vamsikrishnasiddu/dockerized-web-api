package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	fmt.Println("server is running on 8081")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("/ is called.")
		log.Println("/ is called.")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/hi is called.")
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
