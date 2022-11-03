package bgg

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
