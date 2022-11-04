package bgg

import "log"

type Game struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	BggId         string `json:"bggId"`
	BggType       []ThingType
	Thumbnail     string `json:"thumbnail"`
	Image         string `json:"image"`
	Yearpublished string `json:"yearpublished"`
}

type GameCollection map[string]Game

func CreateGCfromSR(sr BggSearchResult) GameCollection {
	gc := GameCollection{}

	for _, v := range sr.Item {
		if e, ok := gc[v.ID]; ok {
			e.BggType = append(e.BggType, getThingType(v.Type))
			gc[v.ID] = e
		} else {
			tt := []ThingType{}
			gc[v.ID] = Game{
				Title:         v.Name.Value,
				BggId:         v.ID,
				BggType:       append(tt, getThingType(v.Type)),
				Yearpublished: v.Yearpublished.Value,
			}
		}

	}

	return gc
}

func CreateGCfromTI(ti ThingItems) GameCollection {
	gc := GameCollection{}

	for _, v := range ti.Item {
		if e, ok := gc[v.ID]; ok {
			e.BggType = append(e.BggType, getThingType(v.Type))
		} else {
			tt := []ThingType{}
			gc[v.ID] = Game{
				Title:         v.Name[0].Value,
				BggId:         v.ID,
				BggType:       append(tt, getThingType(v.Type)),
				Yearpublished: v.Yearpublished.Value,
				Description:   v.Description,
				Thumbnail:     v.Thumbnail,
				Image:         v.Image,
			}
		}
	}

	return gc
}

func (gc *GameCollection) getThings() {
	bggIds := make([]string, len(*gc))
	for k := range *gc {
		bggIds = append(bggIds, k)
	}

	tq := NewThingQuery(bggIds...)
	gc, err := Query(tq)
	if err != nil {
		log.Println(err)
	}
}
