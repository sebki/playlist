package bgg

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/sebki/playlist/internal/models"
)

type HotlistType string

const (
	BoardgameHotlistType        HotlistType = "boardgame"        // BoardgameHotlistType is the type for boardgames
	RPGHotlistType              HotlistType = "rpg"              // RPGHotlistType is the type for rpgs
	VideogameHotlistType        HotlistType = "videogame"        // VideogameHotlistType is the type for videogames
	BoardgamePersonHotlistType  HotlistType = "boardgameperson"  // BoardgamePersonHotlistType is the type for boardgamepersons
	RPGPersonHotlistType        HotlistType = "rpgperson"        // RPGPersonHotlistType is the type for rpgpersons
	BoardgameCompanyHotlistType HotlistType = "boardgamecompany" // BoardgameCompanyHotlistType is the type for boardgamecompanies
	RPGCompanyHotlistType       HotlistType = "rpgcompany"       // RPGCompanyHotlistType is the type for rpgcompanies
	VideogameCompanyHotlistType HotlistType = "videogamecompany" // VideogameCompanyHotlistType is the type for videogamecompanies
)

// HotQuery retrieves the list of most active items on the site.
type HotQuery struct {
	hotlistType HotlistType
}

// NewHotQuery returns a pointer to a new HotQuery
func NewHotQuery(hotness string) *HotQuery {
	hq := HotQuery{
		hotlistType: HotlistType(hotness),
	}
	return &hq
}

func (hq *HotQuery) generateSearchString() string {
	searchString := baseURL + "hot?type=" + string(hq.hotlistType)
	return searchString
}

// HotItems contains all response data from a HotQuery
type BggHotResult struct {
	XMLName    xml.Name `xml:"items"`
	Termsofuse string   `xml:"termsofuse,attr"`
	Item       []struct {
		ID        string `xml:"id,attr"`
		Rank      string `xml:"rank,attr"`
		Thumbnail struct {
			Value string `xml:"value,attr"`
		} `xml:"thumbnail"`
		Name struct {
			Value string `xml:"value,attr"`
		} `xml:"name"`
		Yearpublished struct {
			Value string `xml:"value,attr"`
		} `xml:"yearpublished"`
	} `xml:"item"`
}

// Write unmarshals the response body to HotItems
func (bhr *BggHotResult) Unmarshal(b *http.Response) error {
	defer b.Body.Close()
	body, err := io.ReadAll(b.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, bhr)
	if err != nil {
		return err
	}
	return nil
}

func (bhr *BggHotResult) ToGames() []models.Game {
	ids := []string{}
	for _, v := range bhr.Item {
		ids = append(ids, v.ID)
	}
	tq := NewThingQuery(ids...)
	games, err := Query(tq)
	if err != nil {
		log.Println(err)
	}
	return games
}

func (bhr *BggHotResult) ToPlaylist() models.Playlist {
	games := bhr.ToGames()
	index := map[string]int{}

	for i, v := range games {
		index[v.BggId] = i
	}

	playlist := models.Playlist{ListType: models.HotnessListType}

	for _, v := range bhr.Item {
		lg := models.ListedGame{
			Game: games[index[v.ID]],
		}
		lg.SetRank(v.Rank)
		playlist.AddGames(lg)
	}

	return playlist
}
