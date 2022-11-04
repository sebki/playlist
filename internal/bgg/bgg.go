package bgg

import (
	"log"
	"net/http"
)

// BggQuery interface
type BggQuery interface {
	generateSearchString() string
}

type BggResult interface {
	UnmarshalBody(r *http.Response) error
}

const baseURL = "https://www.boardgamegeek.com/xmlapi2/"

type ThingType string

const (
	TypeBoardGame          ThingType = "boardgame"              // TypeBoardGame is the ThingType for boardgames
	TypeBoardGameExpansion ThingType = "boardgameexpansion"     // TypeBoardGameExpansion is the ThingType for boardgame expansions
	TypeBoardGameAccessory ThingType = "boardgameaccessory"     // TypeBoardGameAccessory is the ThingType for boardgame accessories
	TypeVideoGame          ThingType = "videogame"              // TypeVideoGame is the ThingType for videogames
	TypeRPGItem            ThingType = "rpgitem"                // TypeRPGItem ist the ThingType for rpg items
	TypeRPGIssue           ThingType = "rpgissue"               // TypeRPGIssue is the ThingType for rpg issues (periodicals)
	ThingTypeNotRecogniced ThingType = "thingtypenotrecogniced" // TypeNoType for when ThingType is not recogniced

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
		return ThingTypeNotRecogniced
	}
}

// ItemType can ether be a thing, or a family
type ItemType string

const (
	ThingItem  ItemType = "thing"  // ThingItem is the type for things
	FamilyItem ItemType = "family" // FamilyItem is the type for families
)

// DomainType represents domains on boardgamegeek
type DomainType string

const (
	BoardGameDomain DomainType = "boardgame" // BoardGameDomain is the DomainType for boardgames
	RPGDomain       DomainType = "rpg"       // RPGDomain is the DomainType for rpgs
	VideogameDomain DomainType = "videogame" // VideogameDomain is the DomainType for videogames

)

// SortType contains types for sorting
type SortType string

const (
	UsernameSortType SortType = "username" // UsernameSortType sorts for username
	DateSortType     SortType = "date"     // DateSortType sorts for date
)

// HotlistType represents all valid types for hotness lists
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

type LinkType string

const (
	BoardgameCategoryType       LinkType = "boardgamecategory"
	BoardgameMechanicType       LinkType = "boardgamemechanic"
	BoardgameFamilyType         LinkType = "boardgamefamily"
	BoardgameExpansionType      LinkType = "boardgameexpansion"
	BoardgameDesignerType       LinkType = "boardgamedesigner"
	BoardgameArtistType         LinkType = "boardgameartist"
	BoardgamePublisherType      LinkType = "boardgamepublisher"
	BoardgameImplementationType LinkType = "boardgameimplementation"
	LinkTypeNotRecognised       LinkType = "linktypenotrecognised"
)

func getLinkType(lt string) LinkType {
	switch lt {
	case string(BoardgameCategoryType):
		return BoardgameCategoryType
	case string(BoardgameMechanicType):
		return BoardgameMechanicType
	case string(BoardgameFamilyType):
		return BoardgameFamilyType
	case string(BoardgameExpansionType):
		return BoardgameExpansionType
	case string(BoardgameDesignerType):
		return BoardgameDesignerType
	case string(BoardgameArtistType):
		return BoardgameArtistType
	case string(BoardgamePublisherType):
		return BoardgamePublisherType
	case string(BoardgameImplementationType):
		return BoardgameImplementationType
	default:
		return LinkTypeNotRecognised
	}
}

// Query queries the Boardgamegeek XML API 2 and returns a http.Response.
// Retries 10 times, if response status is not ok
func Query(q BggQuery) (GameCollection, error) {
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
			return GameCollection{}, err
		}
		gc := CreateGCfromSR(*sr)
		return gc, nil
	case *ThingQuery:
		log.Println("ThingQuery type identified")
		ti := &ThingItems{}
		err := ti.UnmarshalBody(res)
		if err != nil {
			return GameCollection{}, err
		}
		gc := CreateGCfromTI(*ti)
		return gc, nil
	default:
		return GameCollection{}, nil
	}
}
