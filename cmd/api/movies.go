package main

import (
	"net/http"
	"time"

	"github.com/hhieunnd/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	data := data.Movie{
		ID:       1,
		CreateAt: time.Now(),
		Title:    "Karatekid",
		Year:     2011,
		Runtime:  120,
		Genres:   []string{"action"},
		Version:  1,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:       1,
		CreateAt: time.Now(),
		Title:    "Karatekid",
		Year:     2011,
		Runtime:  120,
		Genres:   []string{"action"},
		Version:  1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
