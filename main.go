package main

import (
	"fmt"
	"log"
	"net/http"
)

var version = "v1"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website! Version %s", version)
		fmt.Fprintf(w, "I updated the site!, Send change", version)
	})
	fmt.Printf("App is starting, version: %s \n", version)
	log.Fatal(http.ListenAndServe(":8500", nil))
}
