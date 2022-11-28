package main

import (
	"fmt"
	"net/http"
)

// logs the error
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// response with a custom error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// response for 500 internal server error, wraps errorResponse()
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// response for 404 not found error
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested ressource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// response for 405 method not allowed error
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this ressource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
