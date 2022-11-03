package bgg

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

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

func (sr *BggSearchResult) getThumbnails() {
	//TODO: Write function when thing lookup is implemented
}

// Write unmarshals the response body to SearchItems
func (sr *BggSearchResult) Unmarshal(b *http.Response) error {
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

type SearchQuery struct {
	Term      string
	ThingType []ThingType
	Exact     bool
}

func (sq *SearchQuery) generateSearchString() string {
	searchString := baseURL + "search?query=" + strings.ReplaceAll(sq.Term, " ", "+")
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

// NewSearchQuery returns a pointer to a new SearchQuery
func NewSearchQuery(query string) *SearchQuery {
	newQuery := strings.ReplaceAll(query, " ", "+")
	sq := SearchQuery{
		Term: newQuery,
	}
	return &sq
}

// SetThingType returns all items that match query of type ThingType
func (sq *SearchQuery) SetThingType(thingType []ThingType) {
	sq.ThingType = thingType
}

// EnableExact limits results to items that match the query exactly
func (sq *SearchQuery) EnableExact() {
	sq.Exact = true
}
