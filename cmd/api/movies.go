package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create movie handler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show movie handler")
}
