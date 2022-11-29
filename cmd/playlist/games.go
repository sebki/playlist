package main

import (
	"fmt"
	"net/http"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title         string   `json:"title"`
		Description   string   `json:"description"`
		BggID         string   `json:"bgg_id"`
		BggType       []string `json:"bgg_type"`
		Thumbnail     string   `json:"thumbnail"`
		Image         string   `json:"image"`
		Yearpublished string   `json:"yearpublished"`
		Links         []struct {
			LinkType  string `json:"link_type"`
			BggId     string `json:"bgg_id"`
			LinkValue string `json:"link_value"`
			Inbound   bool   `json:"inbound"`
		} `json:"links"`
		Minage      string `json:"minage"`
		Minplayer   string `json:"minplayer"`
		Maxplayer   string `json:"maxplayer"`
		Minplaytime string `json:"minplaytime"`
		Maxplaytime string `json:"maxplaytime"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

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
