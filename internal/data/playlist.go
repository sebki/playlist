package data

import (
	"context"
	"time"
)

type Playlist struct {
	UID          string       `json:"uid"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	DateCreated  time.Time    `json:"datecreated"`
	DateModified time.Time    `json:"datemodified"`
	ListType     string       `json:"listtype"`
	Length       int          `json:"length"`
	Games        []ListedGame `json:"games"`
}

func (pl *Playlist) AddGames(lg ...ListedGame) {
	pl.Games = append(pl.Games, lg...)
	pl.Length = len(pl.Games)
}

type PlaylistService interface {
	FindPlaylistByID(ctx context.Context, bggid string) (*Playlist, error)
	FindPlaylists(ctx context.Context, filter PlaylistFilter) (*[]Playlist, error)
	CreatePlaylist(ctx context.Context, game Playlist) error
	UpdatePlaylist(ctx context.Context, bggid string, upd PlaylistUpdate) (*Playlist, error)
	DeletePlaylist(ctx context.Context, bggid string) error
}

type PlaylistFilter struct {
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	DateCreated  time.Time    `json:"datecreated"`
	DateModified time.Time    `json:"datemodified"`
	ListType     string       `json:"listtype"`
	Length       int          `json:"length"`
	Games        []ListedGame `json:"games"`
}

type PlaylistUpdate struct {
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	DateModified time.Time    `json:"datemodified"`
	ListType     string       `json:"listtype"`
	Length       int          `json:"length"`
	Games        []ListedGame `json:"games"`
}
