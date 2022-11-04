package bgg

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// ThingQuery contains all required Data for a "thing"-search on Boardgamegeek
type ThingQuery struct {
	id             []string
	thingType      []ThingType
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
		thingType: []ThingType{},
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

// SetType sets thingType
// Type Specifies that, regardless of the type of
// thing asked for by id, the results are filtered
// by the THINGTYPE(s) specified. Multiple THINGTYPEs
// can be specified in a comma-delimited list.
func (tq *ThingQuery) SetType(types ...ThingType) {
	ttSlice := []ThingType{}

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
type ThingItems struct {
	XMLName    xml.Name `xml:"items"`
	Text       string   `xml:",chardata"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       []struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description   string `xml:"description"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
		Minplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplayers"`
		Maxplayers struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplayers"`
		Poll []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Title      string `xml:"title,attr"`
			Totalvotes string `xml:"totalvotes,attr"`
			Results    []struct {
				Text       string `xml:",chardata"`
				Numplayers string `xml:"numplayers,attr"`
				Result     []struct {
					Text     string `xml:",chardata"`
					Value    string `xml:"value,attr"`
					Numvotes string `xml:"numvotes,attr"`
					Level    string `xml:"level,attr"`
				} `xml:"result"`
			} `xml:"results"`
		} `xml:"poll"`
		Playingtime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"playingtime"`
		Minplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minplaytime"`
		Maxplaytime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"maxplaytime"`
		Minage struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"minage"`
		Link []struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			ID    string `xml:"id,attr"`
			Value string `xml:"value,attr"`
		} `xml:"link"`
		Videos struct {
			Text  string `xml:",chardata"`
			Total string `xml:"total,attr"`
			Video []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Title    string `xml:"title,attr"`
				Category string `xml:"category,attr"`
				Language string `xml:"language,attr"`
				Link     string `xml:"link,attr"`
				Username string `xml:"username,attr"`
				Userid   string `xml:"userid,attr"`
				Postdate string `xml:"postdate,attr"`
			} `xml:"video"`
		} `xml:"videos"`
		Versions struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				ID        string `xml:"id,attr"`
				Thumbnail string `xml:"thumbnail"`
				Image     string `xml:"image"`
				Link      []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					ID      string `xml:"id,attr"`
					Value   string `xml:"value,attr"`
					Inbound string `xml:"inbound,attr"`
				} `xml:"link"`
				Name struct {
					Text      string `xml:",chardata"`
					Type      string `xml:"type,attr"`
					Sortindex string `xml:"sortindex,attr"`
					Value     string `xml:"value,attr"`
				} `xml:"name"`
				Yearpublished struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"yearpublished"`
				Productcode struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"productcode"`
				Width struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"width"`
				Length struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"length"`
				Depth struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"depth"`
				Weight struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"weight"`
			} `xml:"item"`
		} `xml:"versions"`
		Comments struct {
			Text       string `xml:",chardata"`
			Page       string `xml:"page,attr"`
			Totalitems string `xml:"totalitems,attr"`
			Comment    []struct {
				Text     string `xml:",chardata"`
				Username string `xml:"username,attr"`
				Rating   string `xml:"rating,attr"`
				Value    string `xml:"value,attr"`
			} `xml:"comment"`
		} `xml:"comments"`
		Marketplacelistings struct {
			Text    string `xml:",chardata"`
			Listing []struct {
				Text     string `xml:",chardata"`
				Listdate struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"listdate"`
				Price struct {
					Text     string `xml:",chardata"`
					Currency string `xml:"currency,attr"`
					Value    string `xml:"value,attr"`
				} `xml:"price"`
				Condition struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"condition"`
				Notes struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"notes"`
				Link struct {
					Text  string `xml:",chardata"`
					Href  string `xml:"href,attr"`
					Title string `xml:"title,attr"`
				} `xml:"link"`
			} `xml:"listing"`
		} `xml:"marketplacelistings"`
	} `xml:"item"`
}

// Write unmarshals the response body to ThingItems
func (ti *ThingItems) UnmarshalBody(b *http.Response) error {
	defer b.Body.Close()
	body, err := io.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, ti)
	if err != nil {
		return err
	}
	return nil
}
