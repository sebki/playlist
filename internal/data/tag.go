package data

import "time"

type Tag struct {
	ID         int64     `json:"id"`
	Type       TagType   `json:"type"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	Version    int32     `json:"version"`
}

type TagType string

const (
	Category TagType = "boardgamecategory"
	Mechanic TagType = "boardgamemechanic"
	Family   TagType = "boardgamefamily"
)
