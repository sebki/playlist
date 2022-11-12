package models

import (
	"log"
	"strconv"
)

type Game struct {
	Uid           string      `json:"uid,omitempty"`
	Title         string      `json:"title,omitempty"`
	Description   string      `json:"description,omitempty"`
	BggId         string      `json:"bggId,omitempty"`
	BggType       []ThingType `json:"bggtype,omitempty"`
	Thumbnail     string      `json:"thumbnail,omitempty"`
	Image         string      `json:"image,omitempty"`
	Yearpublished int         `json:"yearpublished,omitempty"`
	Links         []Link      `json:"links,omitempty"`
	Minage        int         `json:"minage,omitempty"`
	Minplayer     int         `json:"minplayer,omitempty"`
	Maxplayer     int         `json:"maxplayer,omitempty"`
	Minplaytime   int         `json:"minplaytime,omitempty"`
	Maxplaytime   int         `json:"maxplaytime,omitempty"`
}

func NewGame() Game {
	return Game{}
}

func (g *Game) SetTitle(title string) {
	if title == "" {
		log.Println("no title provided")
	}
	g.Title = title
}

func (g *Game) SetDescription(desc string) {
	g.Description = desc
}

func (g *Game) SetBggId(id string) {
	g.BggId = id
}

func (g *Game) AddBggType(tt string) {
	g.BggType = append(g.BggType, getThingType(tt))
}

func (g *Game) SetThumbnailLink(link string) {
	g.Thumbnail = link
}

func (g *Game) SetImageLink(link string) {
	g.Image = link
}

func (g *Game) SetYearpublished(year string) {
	intYear, err := strconv.Atoi(year)
	if err != nil {
		log.Println(err)
	}
	g.Yearpublished = intYear
}

func (g *Game) SetLinks(link ...Link) {
	g.Links = append(g.Links, link...)
}

func (g *Game) SetMinage(age string) {
	intAge, err := strconv.Atoi(age)
	if err != nil {
		log.Println(err)
	}
	g.Minage = intAge
}

func (g *Game) SetMinplayer(count string) {
	intCount, err := strconv.Atoi(count)
	if err != nil {
		log.Println(err)
	}
	g.Minplayer = intCount
}

func (g *Game) SetMaxplayer(count string) {
	intCount, err := strconv.Atoi(count)
	if err != nil {
		log.Println(err)
	}
	g.Maxplayer = intCount
}

func (g *Game) SetMinplaytime(time string) {
	intTime, err := strconv.Atoi(time)
	if err != nil {
		log.Println(err)
	}
	g.Minplaytime = intTime
}

func (g *Game) SetMaxplaytime(time string) {
	intTime, err := strconv.Atoi(time)
	if err != nil {
		log.Println(err)
	}
	g.Maxplaytime = intTime
}

type ThingType string

const (
	TypeBoardGame          ThingType = "boardgame"              // TypeBoardGame is the ThingType for boardgames
	TypeBoardGameExpansion ThingType = "boardgameexpansion"     // TypeBoardGameExpansion is the ThingType for boardgame expansions
	TypeBoardGameAccessory ThingType = "boardgameaccessory"     // TypeBoardGameAccessory is the ThingType for boardgame accessories
	TypeVideoGame          ThingType = "videogame"              // TypeVideoGame is the ThingType for videogames
	TypeRPGItem            ThingType = "rpgitem"                // TypeRPGItem ist the ThingType for rpg items
	TypeRPGIssue           ThingType = "rpgissue"               // TypeRPGIssue is the ThingType for rpg issues (periodicals)
	ThingTypeNotRecognised ThingType = "thingtypenotrecognised" // TypeNoType for when ThingType is not recogniced

)

func getThingType(tt string) ThingType {
	switch tt {
	case string(TypeBoardGame):
		return TypeBoardGame
	case string(TypeBoardGameExpansion):
		return TypeBoardGameExpansion
	case string(TypeBoardGameAccessory):
		return TypeBoardGameAccessory
	case string(TypeVideoGame):
		return TypeVideoGame
	case string(TypeRPGItem):
		return TypeRPGItem
	case string(TypeRPGIssue):
		return TypeRPGIssue
	default:
		return ThingTypeNotRecognised
	}
}
