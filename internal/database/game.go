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
		game(func: eq(bggId, %q)) {
			uid
		}
		
	}`, bggId)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return "", err
	}

	var data struct {
		Uid string `json:"uid"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return "", err
	}

	return data.Uid, nil
}

func (db *db) CreateGames(game ...models.Game) error {
	for _, v := range game {
		if uid, err := db.getUidByBggId(v.BggId); uid == "" {
			if err != nil {
				return err
			}
			log.Println("Create new game in dgraph for game: ", v)
			ctx := context.Background()
			txn := db.Client.NewTxn()
			defer txn.Discard(ctx)

			g, err := json.Marshal(&v)
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
			log.Println("Game found: ", uid)
		}

	}
	return nil
}
