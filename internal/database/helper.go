package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (db *db) mutate(item []byte, uid string) (string, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		SetJson:   item,
		CommitNow: true,
	}

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		return "", err
	}
	newUid := []string{}
	for _, v := range assigned.Uids {
		newUid = append(newUid, v)
	}
	// Mutation on existing Node won't return an uid, return supplied uid
	if len(newUid) == 0 {
		return uid, nil
	}

	return newUid[0], nil

}

func (db *db) getUidByBggId(bggId string) (uid string, err error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		items(func: eq(bggid, %q)) {
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
