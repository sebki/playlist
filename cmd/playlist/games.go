package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sebki/playlist/internal/data"
	"github.com/sebki/playlist/internal/validator"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title         string   `json:"title"`
		Description   string   `json:"description"`
		Yearpublished int32    `json:"yearpublished"`
		Type          []string `json:"bgg_type"`
		Thumbnail     string   `json:"thumbnail"`
		Image         string   `json:"image"`
		Minplayer     int32    `json:"minplayer"`
		Maxplayer     int32    `json:"maxplayer"`
		Minplaytime   int32    `json:"minplaytime"`
		Maxplaytime   int32    `json:"maxplaytime"`
		Minage        int32    `json:"minage"`
		Maxage        int32    `json:"maxage"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	game := &data.Game{
		Title:         input.Title,
		Description:   input.Description,
		YearPublished: input.Yearpublished,
		Type:          *data.ConvertSliceRaw(input.Type),
		Thumbnail:     input.Thumbnail,
		Image:         input.Image,
		MinPlayer:     input.Minplayer,
		MaxPlayer:     input.Maxplayer,
		MinPlaytime:   input.Minplaytime,
		MaxPlaytime:   input.Maxplaytime,
		MinAge:        input.Minage,
		MaxAge:        input.Maxage,
	}

	v := validator.New()

	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Games.Insert(game)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// information vor the client, where the newly created resource is to be found
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/games/%d", game.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"game": game}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	movie, err := app.models.Games.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
