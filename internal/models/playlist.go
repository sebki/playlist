package models

import (
	"log"
	"strconv"
	"time"
)

type ListedGame struct {
	UID             string `json:"uid"`
	Rank            int    `json:"rank"`
	UserDescription string `json:"userdescription"`
	Game            Game
}

func (lg *ListedGame) SetRank(rank string) {
	intRank, err := strconv.Atoi(rank)
	if err != nil {
		log.Println(err)
	}
	lg.Rank = intRank
}

type ListType string

const (
	HotnessListType       ListType = "hotness"
	SearchResultListType  ListType = "searchresult"
	ListTypeNotRecognised ListType = "listtypenotrecognised"
)

func getListType(lt string) ListType {
	switch lt {
	case string(HotnessListType):
		return HotnessListType
	case string(SearchResultListType):
		return SearchResultListType
	default:
		return ListTypeNotRecognised
	}
}

type Playlist struct {
	UID              string       `json:"uid"`
	Title            string       `json:"title"`
	Description      string       `json:"description"`
	DateCreated      time.Time    `json:"datecreated"`
	DateLastModified time.Time    `json:"datemodified"`
	ListType         ListType     `json:"listtype"`
	Length           int          `json:"length"`
	Games            []ListedGame `json:"games"`
}

func (pl *Playlist) SetListType(lt string) {
	pl.ListType = getListType(lt)
}

func (pl *Playlist) AddGames(lg ...ListedGame) {
	pl.Games = append(pl.Games, lg...)
	pl.Length = len(pl.Games)
}
