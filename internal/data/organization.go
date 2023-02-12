package data

import "time"

type Organization struct {
	Id         int64            `json:"id"`
	Type       OrganizationType `json:"type"`
	Value      string           `json:"value"`
	CreatedAt  time.Time        `json:"-"`
	ModifiedAt time.Time        `json:"-"`
	Version    int32            `json:"version"`
}

type OrganizationType string

const (
	Publisher OrganizationType = "boardgamepublisher"
)
