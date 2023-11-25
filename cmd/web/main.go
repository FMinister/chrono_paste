package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/chrono/view", chronoView)
	mux.HandleFunc("/chrono/create", chronoCreate)

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

// 2.05-url-query-strings.html
