package data

import "time"

type Person struct {
	Id         int64      `json:"id"`
	Type       PersonType `json:"type"`
	Value      string     `json:"value"`
	CreatedAt  time.Time  `json:"-"`
	ModifiedAt time.Time  `json:"-"`
	Version    int32      `json:"version"`
}

type PersonType string

const (
	Designer PersonType = "boardgamedesigner"
	Artist   PersonType = "boardgameartist"
)
