package database

import (
	"encoding/json"

	"github.com/sebki/playlist/internal/models"
)

func (db *db) MutateGame(game models.Game) (models.Game, error) {
	if game.Uid == "" {
		uid, err := db.getUidByBggId(game.BggId)
		if err != nil {
			return game, err
		}
		game.Uid = uid
	}

	// for i, v := range game.Links {
	// 	l, err := db.MutateLink(v)
	// 	if err != nil {
	// 		return game, err
	// 	}
	// 	game.Links[i] = l

	// }

	g, err := json.Marshal(&game)
	if err != nil {
		return game, err
	}

	uid, err := db.mutate(g, game.Uid)
	if err != nil {
		return game, err
	}

	game.Uid = uid

	return game, nil

}
