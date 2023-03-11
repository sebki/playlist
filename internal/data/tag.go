package data

import "time"

// All descriptors within the app are tags
type Tag struct {
	ID         int64     `json:"id"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	Version    int32     `json:"version"`
}
