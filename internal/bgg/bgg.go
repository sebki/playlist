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
func Query(q BggQuery) (models.GameCollection, error) {
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
		log.Println("SearchQuery type identified")
		sr := &BggSearchResult{}
		err := sr.UnmarshalBody(res)
		if err != nil {
			return models.GameCollection{}, err
		}
		gc := models.CreateGCfromSR(*sr)
		return gc, nil
	case *ThingQuery:
		log.Println("ThingQuery type identified")
		ti := &ThingItems{}
		err := ti.UnmarshalBody(res)
		if err != nil {
			return models.GameCollection{}, err
		}
		gc := models.CreateGCfromTI(*ti)
		return gc, nil
	default:
		return models.GameCollection{}, nil
	}
}
