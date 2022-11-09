package models

import (
	"log"
	"strconv"
	"time"
)

type ListedGame struct {
	UID         string
	Rank        int
	Description string
	Game        Game
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
	UID              string
	DateCreated      time.Time
	DateLastModified time.Time
	ListType         ListType
	Length           int
	Games            []ListedGame
}

func (pl *Playlist) SetListType(lt string) {
	pl.ListType = getListType(lt)
}

func (pl *Playlist) AddGames(lg ...ListedGame) {
	pl.Games = append(pl.Games, lg...)
	pl.Length = len(pl.Games)
}
