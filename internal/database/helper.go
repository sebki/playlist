package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (db *db) BggIdIsExist(id string) bool {
	check, _ := db.getUidByBggId(id)
	return check != ""
}

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
	q := fmt.Sprintf(`
	{
		items(func: eq(bggid, %q)) {
			uid
		}
		
	}`, bggId)

	res, err := db.query(q)
	if err != nil {
		return "", err
	}

	var data struct {
		Items []struct {
			Uid string `json:"uid"`
		} `json:"items"`
	}

	err = json.Unmarshal(res, &data)
	if err != nil {
		return "", err
	}

	if len(data.Items) > 0 {
		return data.Items[0].Uid, nil
	}

	return "", nil
}

func (db *db) query(query string) ([]byte, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	resp, err := txn.Query(ctx, query)
	if err != nil {
		return []byte{}, err
	}

	return resp.GetJson(), nil
}
