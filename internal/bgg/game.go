package bgg

import "log"

type Links struct {
	Type    LinkType `json:"type"`
	ID      string   `json:"id"`
	Value   string   `json:"value"`
	Inbound bool     `json:"inbound,omitempty"`
}

type Game struct {
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	BggId         string      `json:"bggId"`
	BggType       []ThingType `json:"type"`
	Thumbnail     string      `json:"thumbnail"`
	Image         string      `json:"image"`
	Yearpublished string      `json:"yearpublished"`
	Links         []Links     `json:"links"`
	Minage        string      `json:"minage"`
	Minplayer     string      `json:"minplayer"`
	Maxplayer     string      `json:"maxplayer"`
	Minplaytime   string      `json:"minplaytime"`
	Maxplaytime   string      `json:"maxplaytime"`
}

type GameCollection map[string]Game

func CreateGCfromSR(sr BggSearchResult) GameCollection {
	gc := GameCollection{}

	for _, v := range sr.Item {
		if e, ok := gc[v.ID]; ok {
			e.BggType = append(e.BggType, getThingType(v.Type))
			gc[v.ID] = e
		} else {
			tq := NewThingQuery(v.ID)
			game, err := Query(tq)
			if err != nil {
				log.Println(err)
			}
			if el, io := game[v.ID]; io {
				gc[v.ID] = el
			}
		}
	}

	return gc
}

func CreateGCfromTI(ti ThingItems) GameCollection {
	gc := GameCollection{}

	for _, v := range ti.Games {
		if e, ok := gc[v.ID]; ok {
			e.BggType = append(e.BggType, getThingType(v.Type))
		} else {
			tt := []ThingType{}
			lx := []Links{}
			for _, l := range v.Link {
				newLx := Links{
					Type:    getLinkType(l.Type),
					Value:   l.Value,
					ID:      l.ID,
					Inbound: l.Inbound,
				}
				lx = append(lx, newLx)
			}
			gc[v.ID] = Game{
				Title:         v.Name[0].Value,
				BggId:         v.ID,
				BggType:       append(tt, getThingType(v.Type)),
				Yearpublished: v.Yearpublished.Value,
				Description:   v.Description,
				Thumbnail:     v.Thumbnail,
				Image:         v.Image,
				Minage:        v.Minage.Value,
				Minplayer:     v.Minplayers.Value,
				Maxplayer:     v.Maxplayers.Value,
				Minplaytime:   v.Minplaytime.Value,
				Maxplaytime:   v.Minplaytime.Value,
				Links:         lx,
			}
		}
	}

	return gc
}

func (gc *GameCollection) Array() (games []Game) {
	for _, v := range *gc {
		games = append(games, v)
	}

	return games
}
