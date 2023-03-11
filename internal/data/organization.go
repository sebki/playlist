package data

import "time"

type Organization struct {
	Id         int64     `json:"id"`
	Value      string    `json:"value"`
	Tags       []Tag     `json:"tags"`
	Persons    []Person  `json:"persons"`
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	Version    int32     `json:"version"`
}
