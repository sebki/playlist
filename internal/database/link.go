package database

import (
	"encoding/json"

	"github.com/sebki/playlist/internal/models"
)

func (db *db) MutateLink(link models.Link) (models.Link, error) {
	if link.UID == "" {
		uid, err := db.getUidByBggId(link.BggId)
		if err != nil {
			return link, err
		}

		link.UID = uid
	}

	l, err := json.Marshal(&link)
	if err != nil {
		return link, err
	}

	uid, err := db.mutate(l, link.UID)
	if err != nil {
		return link, err
	}

	link.UID = uid
	return link, nil
}
