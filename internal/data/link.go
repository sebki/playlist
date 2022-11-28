package data

import (
	"context"
	"time"
)

type Link struct {
	UID          string    `json:"uid,omitempty"`
	LinkType     string    `json:"linktype,omitempty"`
	BggId        string    `json:"bggid,omitempty"`
	LinkValue    string    `json:"linkvalue,omitempty"`
	Inbound      bool      `json:"inbound,omitempty"`
	DgraphType   []string  `json:"dgraph.type"`
	LastBggQuery time.Time `json:"lastbggquery"`
}

type LinkService interface {
	FindLinkByID(ctx context.Context, bggid string) (*Link, error)
	FindLinks(ctx context.Context, filter LinkFilter) (*[]Link, error)
	CreateLink(ctx context.Context, link Link) error
	UpdateLink(ctx context.Context, bggid string, upd LinkUpdate) (*Link, error)
	DeleteLink(ctx context.Context, bggid string) error
}

type LinkFilter struct {
	LinkType     string    `json:"linktype,omitempty"`
	LinkValue    string    `json:"linkvalue,omitempty"`
	LastBggQuery time.Time `json:"lastbggquery"`
}

type LinkUpdate struct {
	LastBggQuery time.Time `json:"lastbggquery"`
}
