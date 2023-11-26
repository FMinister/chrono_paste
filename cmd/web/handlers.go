package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) chronoView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		http.NotFound(w, r)
		return
	}
	if id < 1 {
		app.logger.Error("id is < 1", slog.Int("id", id), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific chrono with ID %d...", id)
}

func (app *application) chronoCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		app.logger.Info("method not allowed", slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new chrono..."))
}
