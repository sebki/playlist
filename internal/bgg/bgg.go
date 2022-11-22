package bgg

import (
	"log"
	"net/http"

	"github.com/sebki/playlist/internal/models"
)

// BggQuery interface
type BggQuery interface {
	generateSearchString() string
}

type BggResult interface {
	UnmarshalBody(r *http.Response) error
}

const baseURL = "https://www.boardgamegeek.com/xmlapi2/"

// Query queries the Boardgamegeek XML API 2 and returns a http.Response.
// Retries 10 times, if response status is not ok
func Query(q BggQuery) ([]models.Game, error) {
	log.Println("Query func called")
	search := q.generateSearchString()
	log.Println("Searchstring generated: ", search)

	res, err := http.Get(search)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("BGG get func called and gotten response: ", res.Body)

	switch q.(type) {
	case *SearchQuery:
		sr := &BggSearchResult{}
		err := sr.UnmarshalBody(res)
		if err != nil {
			return []models.Game{}, err
		}
		return sr.ToGames(), nil
	case *ThingQuery:
		btr := &BggThingResult{}
		err := btr.UnmarshalBody(res)
		if err != nil {
			return []models.Game{}, err
		}
		return btr.ToGames(), nil
	case *HotQuery:
		hqr := &BggHotResult{}
		err := hqr.UnmarshalBody(res)
		if err != nil {
			return []models.Game{}, err
		}
		return hqr.ToGames(), nil
	case *FamilyQuery:
		bfr := &BggFamilyResult{}
		err := bfr.UnmarshalBody(res)
		if err != nil {
			return []models.Game{}, err
		}
		return bfr.ToGames(), nil
	default:
		return []models.Game{}, nil
	}
}
