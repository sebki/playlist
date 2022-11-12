package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sebki/playlist/internal/models"
)

func (db *db) searchGameByBggId(bggId string) (models.Game, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		game(func: eq(bggId, %q)) {
			uid
			bggId
			description
			bggId
			bggType {
				
			}

		}
		
	}`, bggId)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return models.Game{}, err
	}

	var data struct {
		models.Game
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return models.Game{}, err
	}

	return data.Game, nil
}
