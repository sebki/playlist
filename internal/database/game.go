package database

import "github.com/sebki/playlist/internal/bgg"

type Game struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	BggId         string `json:"bggId"`
	BggType       []bgg.ThingType
	Thumbnail     string `json:"thumbnail"`
	Image         string `json:"image"`
	Yearpublished string `json:"yearpublished"`
}

type GameCollection map[string]Game

func CreateGCfromSR(sr bgg.SearchResult) GameCollection {
	gc := GameCollection{}

	for _, v := range sr.Item {
		if _, ok := gc[v.ID]; ok {
			gc[v.ID].BggType = append(gc[v.ID].BggType, bgg.ThingType(v.Type))
		} else {
			gc[v.ID] = Game{
				Title:         v.Name.Value,
				BggId:         v.ID,
				BggType:       v.Type,
				Yearpublished: v.Yearpublished.Value,
			}
		}

	}
}
