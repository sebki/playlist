package bgg

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/sebki/playlist/internal/models"
)

type SearchQuery struct {
	Term       string
	ThingTypes []string
	Exact      bool
}

func (sq *SearchQuery) generateSearchString() string {
	searchString := baseURL + "search?query=" + strings.ReplaceAll(sq.Term, " ", "+")
	if len(sq.ThingTypes) > 0 {
		searchString += "&type="
		for i, v := range sq.ThingTypes {
			searchString += v
			if i < len(sq.ThingTypes)-1 {
				searchString += ","
			}
		}

	}
	if sq.Exact {
		searchString += "&exact=1"
	}
	return searchString
}

// NewSearchQuery returns a pointer to a new SearchQuery
func NewSearchQuery(query string) *SearchQuery {
	newQuery := strings.ReplaceAll(query, " ", "+")
	sq := SearchQuery{
		Term: newQuery,
	}
	return &sq
}

// SetThingType returns all items that match query of type ThingType
func (sq *SearchQuery) AddThingType(thingType ...string) {
	sq.ThingTypes = append(sq.ThingTypes, thingType...)
}

// EnableExact limits results to items that match the query exactly
func (sq *SearchQuery) EnableExact() {
	sq.Exact = true
}

type BggSearchResult struct {
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

func (bsr *BggSearchResult) ToGames() []models.Game {
	ids := []string{}
	for _, v := range bsr.Item {
		ids = append(ids, v.ID)
	}
	tq := NewThingQuery(ids...)
	games, err := Query(tq)
	if err != nil {
		log.Println(err)
	}

	return games
}

// UnmarshalBody wraps xml.Unmarshal
func (sr *BggSearchResult) UnmarshalBody(b *http.Response) error {
	defer b.Body.Close()
	body, err := io.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, sr)
	if err != nil {
		return err
	}
	return nil
}
