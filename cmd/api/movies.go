package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hhieunnd/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title" validate:"required,max=500"`
		Year    int32        `json:"year" validate:"gte=1888"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	errorValidation := app.validate(input)

	if errorValidation != nil {
		app.failedValidationResponse(w, r, errorValidation)
		return
	}

	fmt.Fprintf(w, "got inputs %+v\n", input)
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
