package main

import (
	"fmt"
	"net/http"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Game created")
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	// TODO: Database-lookup with provided id, JSON response
	fmt.Fprintf(w, "Game with id %d will be shown", id)
}
