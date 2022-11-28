package data

import (
	"context"
	"log"
	"strconv"
)

type ListedGame struct {
	UID             string `json:"uid"`
	Rank            int    `json:"rank"`
	UserDescription string `json:"userdescription"`
	Game            *Game
}

func (lg *ListedGame) SetRank(rank string) {
	intRank, err := strconv.Atoi(rank)
	if err != nil {
		log.Println(err)
	}
	lg.Rank = intRank
}

type ListedGameService interface {
	FindListedGameByID(ctx context.Context, bggid string) (*ListedGame, error)
	FindListedGames(ctx context.Context, filter ListedGameFilter) (*[]ListedGame, error)
	CreateListedGame(ctx context.Context, game ListedGame) error
	UpdateListedGame(ctx context.Context, bggid string, upd ListedGameUpdate) (*ListedGame, error)
	DeleteListedGame(ctx context.Context, bggid string) error
}

type ListedGameFilter struct {
	Rank            int    `json:"rank"`
	UserDescription string `json:"userdescription"`
}

type ListedGameUpdate struct {
	Rank            int    `json:"rank"`
	UserDescription string `json:"userdescription"`
}
