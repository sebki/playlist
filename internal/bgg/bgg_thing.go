package bgg

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/sebki/playlist/internal/models"
)

// ThingQuery contains all required Data for a "thing"-search on Boardgamegeek
type ThingQuery struct {
	id             []string
	thingType      []string
	versions       bool
	videos         bool
	stats          bool
	historical     bool
	marketplace    bool
	comments       bool
	ratingComments bool
	page           int
	pageSize       int
}

// NewThingQuery generates a new ThingQuery with the provided ids
// ID Specifies the id of the thing(s) to retrieve.
func NewThingQuery(ids ...string) *ThingQuery {
	idSlice := []string{}
	idSlice = append(idSlice, ids...)
	tq := ThingQuery{
		id:        idSlice,
		thingType: []string{},
	}
	return &tq
}

// generateSearchString generates a search URL from data provided in
// ThingQuery, fulfills the BggQuery interfaces
func (tq *ThingQuery) generateSearchString() string {
	searchString := baseURL + "thing?"
	if len(tq.id) <= 0 {
		return ""
	}
	searchString += "id="
	for i, id := range tq.id {
		if i+1 > 1 {
			searchString += ","
		}
		searchString += id
	}
	if len(tq.thingType) > 0 {
		searchString += "&"
		for i, t := range tq.thingType {
			if i+1 > 1 {
				searchString += ","
			}
			searchString += string(t)
		}
	}
	if tq.versions {
		searchString += "&versions=1"
	}
	if tq.videos {
		searchString += "&videos=1"
	}
	if tq.stats {
		searchString += "&stats=1"
	}
	if tq.historical {
		searchString += "&historical=1"
	}
	if tq.marketplace {
		searchString += "&marketplace=1"
	}
	if tq.comments {
		searchString += "&comments=1"
	}
	if tq.ratingComments {
		searchString += "&ratingcomments=1"
	}
	if tq.page >= 10 && tq.page <= 100 {
		pageNumber := strconv.Itoa(tq.page)
		searchString += "&page=" + pageNumber
	}
	if tq.pageSize > 0 {
		sizeNumber := strconv.Itoa(tq.pageSize)
		searchString += "&pagesize" + sizeNumber
	}
	return searchString
}

// AddType sets thingType
// Type Specifies that, regardless of the type of
// thing asked for by id, the results are filtered
// by the THINGTYPE(s) specified. Multiple THINGTYPEs
// can be specified in a comma-delimited list.
func (tq *ThingQuery) AddType(types ...string) {
	ttSlice := []string{}

	ttSlice = append(ttSlice, types...)

	tq.thingType = ttSlice
}

// EnableVersions sets versions to true
// returns version info for the item.
func (tq *ThingQuery) EnableVersions() {
	tq.versions = true
}

// EnableVideos sets videos to true
func (tq *ThingQuery) EnableVideos() {
	tq.videos = true
}

// EnableStats sets stats to true
// returns ranking and rating stats for the item.
func (tq *ThingQuery) EnableStats() {
	tq.stats = true
}

// EnableHistorical sets historical to true
// returns historical data over time. See page parameter.
func (tq *ThingQuery) EnableHistorical() {
	tq.historical = true
}

// EnableMarketplace sets marketplace to true
// returns marketplace data.
func (tq *ThingQuery) EnableMarketplace() {
	tq.marketplace = true
}

// EnableComments sets comments to true
// If true, returns all comments about the item. Also includes
// ratings when commented. See page parameter.
func (tq *ThingQuery) EnableComments() {
	tq.comments = true
}

// EnableRatingcomments sets ratingcomments to true
// returns all ratings for the item. Also includes comments
// when rated. See page parameter. The ratingcomments and comments
// parameters cannot be used together, as the output always appears
// in the <comments> node of the XML; comments parameter takes
// precedence if both are specified. Ratings are sorted in descending
// rating value, based on the highest rating they have assigned to that
// item (each item in the collection can have a different rating).
func (tq *ThingQuery) EnableRatingcomments() error {
	if tq.comments {
		return errors.New("comments already enabled")
	}
	tq.ratingComments = true
	return nil
}

// SetPage sets page to the provided value
// Defaults to 1, controls the page of data to see for historical info,
// comments, and ratings data.
func (tq *ThingQuery) SetPage(p int) {
	tq.page = p
}

// SetPagesize sets the number of records to return in paging. Minimum is 10,
// maximum is 100.
func (tq *ThingQuery) SetPagesize(ps int) error {
	if ps < 10 || ps > 100 {
		return errors.New("value must be between 10 and 100")
	}
	tq.pageSize = ps
	return nil
}

// ThingItems contains all possible data response of a "thing"-query on Boardgamegeek
type BggThingResult struct {
	XMLName    xml.Name `xml:"items"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Games      []struct {
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description   string `xml:"description"`
		Yearpublished struct {
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
		Minplayers struct {
			Value string `xml:"value,attr"`
		} `xml:"minplayers"`
		Maxplayers struct {
			Value string `xml:"value,attr"`
		} `xml:"maxplayers"`
		Playingtime struct {
			Value string `xml:"value,attr"`
		} `xml:"playingtime"`
		Minplaytime struct {
			Value string `xml:"value,attr"`
		} `xml:"minplaytime"`
		Maxplaytime struct {
			Value string `xml:"value,attr"`
		} `xml:"maxplaytime"`
		Minage struct {
			Value string `xml:"value,attr"`
		} `xml:"minage"`
		Link []struct {
			Type    string `xml:"type,attr"`
			ID      string `xml:"id,attr"`
			Value   string `xml:"value,attr"`
			Inbound bool   `xml:"inbound,attr"`
		} `xml:"link"`
	} `xml:"item"`
}

// Write unmarshals the response body to ThingItems
func (btr *BggThingResult) UnmarshalBody(b *http.Response) error {
	defer b.Body.Close()
	body, err := io.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, btr)
	if err != nil {
		return err
	}
	return nil
}

func (btr *BggThingResult) ToGameCollection() models.GameCollection {
	gc := models.GameCollection{}

	for _, v := range btr.Games {
		if e, ok := gc[v.ID]; ok {
			e.SetBggType(v.Type)
		} else {
			lx := []models.Link{}
			for _, l := range v.Link {

				newLx := models.Link{
					Value:   l.Value,
					ID:      l.ID,
					Inbound: l.Inbound,
				}
				newLx.SetLinkType(v.Type)
				lx = append(lx, newLx)
			}
			game := models.Game{
				Title:       v.Name[0].Value,
				BggId:       v.ID,
				Description: v.Description,
				Thumbnail:   v.Thumbnail,
				Image:       v.Image,
				Links:       lx,
			}
			game.SetBggType(v.Type)
			game.SetYearpublished(v.Yearpublished.Value)
			game.SetMinage(v.Minage.Value)
			game.SetMinplaytime(v.Minplaytime.Value)
			game.SetMaxplaytime(v.Maxplaytime.Value)
			game.SetMinplayer(v.Minplayers.Value)
			game.SetMaxplayer(v.Maxplayers.Value)

		}
	}

	return gc
}
