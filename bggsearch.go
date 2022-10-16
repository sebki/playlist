package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const BaseURL = "https://www.boardgamegeek.com/xmlapi2/"

type ThingType string

const (
	TypeBoardGame          ThingType = "boardgame"          // TypeBoardGame is the ThingType for boardgames
	TypeBoardGameExpansion ThingType = "boardgameexpansion" // TypeBoardGameExpansion is the ThingType for boardgame expansions
	TypeBoardGameAccessory ThingType = "boardgameaccessory" // TypeBoardGameAccessory is the ThingType for boardgame accessories
	TypeVideoGame          ThingType = "videogame"          // TypeVideoGame is the ThingType for videogames
	TypeRPGItem            ThingType = "rpgitem"            // TypeRPGItem ist the ThingType for rpg items
	TypeRPGIssue           ThingType = "rpgissue"           // TypeRPGIssue is the ThingType for rpg issues (periodicals)

)

type SearchResult struct {
	Total string `xml:"total,attr" json:"total"`
	Item  []struct {
		Type      string `xml:"type,attr" json:"type"`
		ID        string `xml:"id,attr" json:"id"`
		Thumbnail string `json:"thumbnail"`
		Name      struct {
			Type  string `xml:"type,attr" json:"type"`
			Value string `xml:"value,attr" json:"value"`
		} `xml:"name" json:"name"`
		Yearpublished struct {
			Value string `xml:"value,attr" json:"value"`
		} `xml:"yearpublished" json:"yearpublished"`
	} `xml:"item" json:"item"`
}

func (sr *SearchResult) getThumbnails() {
	//TODO: Write function when thing lookup is implemented
}

type SearchQuery struct {
	Term      string
	ThingType []ThingType
	Exact     bool
}

func (sq *SearchQuery) generateSearchString() string {
	searchString := BaseURL + "search?query=" + strings.ReplaceAll(sq.Term, " ", "+")
	if len(sq.ThingType) > 0 {
		searchString += "&type="
		for i, v := range sq.ThingType {
			searchString += string(v)
			if i < len(sq.ThingType)-1 {
				searchString += ","
			}
		}

	}
	if sq.Exact {
		searchString += "&exact=1"
	}
	return searchString
}

func bggsearch(c *gin.Context) {
	res := SearchResult{}
	search := SearchQuery{
		ThingType: []ThingType{
			TypeBoardGame,
			TypeBoardGameExpansion,
		},
		Exact: false,
	}

	searchTerm, isExist := c.GetQuery("query")
	if !isExist {
		c.JSON(200, res)
		return
	}

	search.Term = searchTerm

	searchString := search.generateSearchString()

	httpRes, err := http.Get(searchString)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	defer httpRes.Body.Close()
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	err = xml.Unmarshal(body, &res)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(200, res)
}
