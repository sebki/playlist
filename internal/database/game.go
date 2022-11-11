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
		return err
	}

	var data struct {
		Username []struct {
			Name string `json:"username"`
		} `json:"user"`
		Email []struct {
			Email string `json:"email"`
		} `json:"email"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return err
	}

}
