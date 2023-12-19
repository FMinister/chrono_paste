package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/FMinister/chrono_paste/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	chronos, err := app.chronos.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Chronos = chronos

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) chronoView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	chrono, err := app.chronos.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Chrono = chrono

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) chronoCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "This is a dummy title"
	content := "This is \n a dummy content"
	expires := 7

	id, err := app.chronos.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/chrono/view?id=%d", id), http.StatusSeeOther)
}
