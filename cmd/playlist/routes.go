package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/games", app.listGamesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/games", app.createGameHandler)
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.showGameHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/games/:id", app.updateGameHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/games/:id", app.deleteGameHandler)

	return app.recoverPanic(app.rateLimit(router))
}
