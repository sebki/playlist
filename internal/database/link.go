package database

import (
	"encoding/json"
	"fmt"
	"time"

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

func (db *db) GetLink(bggid string) (models.Link, error) {
	query := fmt.Sprintf(`
	{
		links(func: eq(bggid, %q)){
			uid
			linktype
			bggid
			linkvalue
			inbound
			dgraph.type
			lastbggquery
		}
	}	
	`, bggid)

	var data struct {
		Links []models.Link `json:"links"`
	}

	res, err := db.query(query)
	if err != nil {
		return models.Link{}, err
	}

	err = json.Unmarshal(res, &data)
	if err != nil {
		return models.Link{}, err
	}
	if len(data.Links) == 0 {
		return models.Link{}, nil
	}
	return data.Links[0], nil
}

func (db *db) GetFamilyLinks(lastQuery time.Time) ([]models.Link, error) {
	date := lastQuery.Format("2006-01-02")
	query := fmt.Sprintf(`
	{
		links(func: eq(linktype, "boardgamefamily")) @filter(lt(lastbggquery, %q)){
			uid
			linktype
			bggid
			linkvalue
			inbound
			dgraph.type
			lastbggquery
		}
	}
	`, date)

	var data struct {
		Links []models.Link `json:"links"`
	}

	res, err := db.query(query)
	if err != nil {
		return data.Links, err
	}

	err = json.Unmarshal(res, &data)
	if err != nil {
		return data.Links, err
	}

	return data.Links, nil
}
