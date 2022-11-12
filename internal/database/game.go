package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/sebki/playlist/internal/models"
)

func (db *db) getUidByBggId(bggId string) (uid string, err error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		items(func: eq(bggId, %q)) {
			uid
		}
		
	}`, bggId)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return "", err
	}

	var data struct {
		Items []struct {
			Uid string `json:"uid"`
		} `json:"items"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return "", err
	}

	if len(data.Items) > 0 {
		return data.Items[0].Uid, nil
	}

	return "", nil
}

func (db *db) CreateGames(games ...models.Game) error {
	for _, game := range games {
		if uid, err := db.getUidByBggId(game.BggID()); uid == "" {
			if err != nil {
				return err
			}
			log.Println("Create new game in dgraph for game: ", game)
			ctx := context.Background()
			txn := db.Client.NewTxn()
			defer txn.Discard(ctx)

			g, err := json.Marshal(&game)
			if err != nil {
				return err
			}

			mu := &api.Mutation{
				SetJson:   g,
				CommitNow: true,
			}

			_, err = txn.Mutate(ctx, mu)
			if err != nil {
				return err
			}

		} else {
			log.Println("Item found: ", uid)
		}
	}

	return nil
}
