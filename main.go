package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from ChronoPaste!"))
}

func chronoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func chronoCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

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
