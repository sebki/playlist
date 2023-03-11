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
		Yearpublished int32    `json:"year_published"`
		Type          []string `json:"bg_type"`
		Thumbnail     string   `json:"thumbnail"`
		Image         string   `json:"image"`
		MinPlayer     int32    `json:"minplayer"`
		MaxPlayer     int32    `json:"maxplayer"`
		MinPlaytime   int32    `json:"minplaytime"`
		MaxPlaytime   int32    `json:"maxplaytime"`
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
		Thumbnail:     input.Thumbnail,
		Image:         input.Image,
		MinPlayer:     input.MinPlayer,
		MaxPlayer:     input.MaxPlayer,
		MinPlaytime:   input.MinPlaytime,
		MaxPlaytime:   input.MaxPlaytime,
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

func (app *application) updateGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	game, err := app.models.Games.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Title, Year and Runtime as pointers, so check for nil-value is possible,
	// necessary for partial updates. If no value for a field is provided, no new
	// value will be assigned to the game struct, see if statements below.
	var input struct {
		Title         *string `json:"title"`
		Description   *string `json:"description"`
		YearPublished *int32  `json:"year_published"`
		Thumbnail     *string `json:"thumbnail"`
		Image         *string `json:"image"`
		MinPlayer     *int32  `json:"minplayer"`
		MaxPlayer     *int32  `json:"maxplayer"`
		MinPlaytime   *int32  `json:"minplaytime"`
		MaxPlaytime   *int32  `json:"maxplaytime"`
		MinAge        *int32  `json:"minage"`
		MaxAge        *int32  `json:"maxage"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		game.Title = *input.Title
	}

	if input.Description != nil {
		game.Description = *input.Description
	}

	if input.YearPublished != nil {
		game.YearPublished = *input.YearPublished
	}

	if input.Thumbnail != nil {
		game.Thumbnail = *input.Thumbnail
	}

	if input.Image != nil {
		game.Image = *input.Image
	}

	if input.MinPlayer != nil {
		game.MinPlayer = *input.MinPlayer
	}

	if input.MaxPlayer != nil {
		game.MaxPlayer = *input.MaxPlayer
	}

	if input.MinPlaytime != nil {
		game.MinPlaytime = *input.MinPlaytime
	}

	if input.MaxPlaytime != nil {
		game.MaxPlaytime = *input.MaxPlaytime
	}

	if input.MinAge != nil {
		game.MinAge = *input.MinAge
	}

	if input.MaxAge != nil {
		game.MaxAge = *input.MaxAge
	}

	v := validator.New()

	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Games.Update(game)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Games.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "game succesfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listGamesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Type  []string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Type = app.readCSV(qs, "bg_type", []string{})
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year_published", "max_playtime", "-id", "-title", "-year_published", "-max_playtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	games, metadata, err := app.models.Games.GetAll(input.Title, input.Type, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"games": games, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
