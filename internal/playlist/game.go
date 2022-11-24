package playlist

import "context"

type Game struct {
	Uid           string   `json:"uid,omitempty"`
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	BggId         string   `jsons:"bggid,omitempty"`
	BggType       []string `json:"bggtype,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty"`
	Image         string   `json:"image,omitempty"`
	Yearpublished string   `json:"yearpublished,omitempty"`
	Links         []Link   `json:"links,omitempty"`
	Minage        string   `json:"minage,omitempty"`
	Minplayer     string   `json:"minplayer,omitempty"`
	Maxplayer     string   `json:"maxplayer,omitempty"`
	Minplaytime   string   `json:"minplaytime,omitempty"`
	Maxplaytime   string   `json:"maxplaytime,omitempty"`
	DgraphType    []string `json:"dgraph.type"`
}

type GameService interface {
	FindGameByID(ctx context.Context, bggid string) (*Game, error)
	FindGames(ctx context.Context, filter GameFilter) (*[]Game, error)
	CreateGame(ctx context.Context, game Game) error
	UpdateGame(ctx context.Context, bggid string, upd GameUpdate) (*Game, error)
	DeleteGame(ctx context.Context, bggid string) error
}

type GameFilter struct {
	Title         string   `json:"title,omitempty"`
	BggType       []string `json:"bggtype,omitempty"`
	Yearpublished string   `json:"yearpublished,omitempty"`
	Links         []Link   `json:"links,omitempty"`
	Minage        string   `json:"minage,omitempty"`
	Minplayer     string   `json:"minplayer,omitempty"`
	Maxplayer     string   `json:"maxplayer,omitempty"`
	Minplaytime   string   `json:"minplaytime,omitempty"`
	Maxplaytime   string   `json:"maxplaytime,omitempty"`
}

type GameUpdate struct {
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Image       string `json:"image,omitempty"`
}
