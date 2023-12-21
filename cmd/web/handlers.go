package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/FMinister/chrono_paste/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
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
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) chronoCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 14

	id, err := app.chronos.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/chrono/view/%d", id), http.StatusSeeOther)
}
