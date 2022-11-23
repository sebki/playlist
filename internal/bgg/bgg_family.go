package bgg

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/sebki/playlist/internal/database"
	"github.com/sebki/playlist/internal/models"
)

// FamilyQuery comntains all data for a family query on Boardgamegeek
type FamilyQuery struct {
	id []string
}

// NewFamilyQuery returns a pointer to a new Familyquery
func NewFamilyQuery(ids ...string) *FamilyQuery {
	idSlice := []string{}
	idSlice = append(idSlice, ids...)

	fq := FamilyQuery{
		id: idSlice,
	}
	return &fq
}

func (fq *FamilyQuery) generateSearchString() string {
	searchString := baseURL + "family?"
	searchString += "id="
	for i, id := range fq.id {
		if i+1 > 1 {
			searchString += ","
		}
		searchString += id
	}
	return searchString
}

// FamilyItems contains bgg response from a family query
type BggFamilyResult struct {
	XMLName    xml.Name `xml:"items"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       struct {
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description string `xml:"description"`
		Link        []struct {
			Type    string `xml:"type,attr"`
			ID      string `xml:"id,attr"`
			Value   string `xml:"value,attr"`
			Inbound string `xml:"inbound,attr"`
		} `xml:"link"`
	} `xml:"item"`
}

// Write unmarshals the response body to FamilyItems
func (bfr *BggFamilyResult) UnmarshalBody(b *http.Response) error {
	defer b.Body.Close()
	body, err := io.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, bfr)
	if err != nil {
		return err
	}
	return nil
}

func (bfr *BggFamilyResult) ToGames() []models.Game {
	ids := []string{}
	for _, v := range bfr.Item.Link {
		if !database.Database.BggIdIsExist(v.ID) {
			ids = append(ids, v.ID)
		}
	}
	if len(ids) == 0 {
		return []models.Game{}
	}

	tq := NewThingQuery(ids...)
	games, err := Query(tq)
	if err != nil {
		log.Println(err)
	}
	return games
}
