package data

import "time"

type Person struct {
	Id         int64     `json:"id"`
	Value      string    `json:"value"`
	Tags       []Tag     `json:"tags"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	Version    int32     `json:"version"`
}
