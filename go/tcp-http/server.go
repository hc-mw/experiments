package main

import (
	"fmt"
	"log"
	"net/http"
)

func main2() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
