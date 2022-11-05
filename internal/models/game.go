package models

type Game struct {
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	BggId         string      `json:"bggId"`
	BggType       []ThingType `json:"type"`
	Thumbnail     string      `json:"thumbnail"`
	Image         string      `json:"image"`
	Yearpublished int         `json:"yearpublished"`
	Links         []Link      `json:"links"`
	Minage        int         `json:"minage"`
	Minplayer     int         `json:"minplayer"`
	Maxplayer     int         `json:"maxplayer"`
	Minplaytime   int         `json:"minplaytime"`
	Maxplaytime   int         `json:"maxplaytime"`
}

func NewGame() Game {
	return Game{}
}

func (g *Game) SetTitle(title string) {
	g.Title = title
}

func (g *Game) SetDescription(desc string) {
	g.Description = desc
}

func (g *Game) SetBggId(id string) {
	g.BggId = id
}

func (g *Game) SetBggType(tt string) {
	g.BggType = append(g.BggType, getThingType(tt))
}

func (g *Game) SetThumbnailLink(link string) {
	g.Thumbnail = link
}

func (g *Game) SetImageLink(link string) {
	g.Image = link
}

func (g *Game) SetYearpublished(year int) {
	g.Yearpublished = year
}

func (g *Game) SetLink(link Link) {
	g.Links = append(g.Links, link)
}

func (g *Game) SetMinage(age int) {
	g.Minage = age
}

func (g *Game) SetMinplayer(count int) {
	g.Minplayer = count
}

func (g *Game) SetMaxplayer(count int) {
	g.Maxplayer = count
}

func (g *Game) SetMinplaytime(time int) {
	g.Minplaytime = time
}

func (g *Game) SetMaxplaytime(time int) {
	g.Maxplaytime = time
}

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

type GameCollection map[string]Game

// func CreateGCfromSR(sr BggSearchResult) GameCollection {
// 	gc := GameCollection{}

// 	for _, v := range sr.Item {
// 		if e, ok := gc[v.ID]; ok {
// 			e.BggType = append(e.BggType, getThingType(v.Type))
// 			gc[v.ID] = e
// 		} else {
// 			tq := NewThingQuery(v.ID)
// 			game, err := Query(tq)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			if el, io := game[v.ID]; io {
// 				gc[v.ID] = el
// 			}
// 		}
// 	}

// 	return gc
// }

// func CreateGCfromTI(ti ThingItems) GameCollection {
// 	gc := GameCollection{}

// 	for _, v := range ti.Games {
// 		if e, ok := gc[v.ID]; ok {
// 			e.BggType = append(e.BggType, getThingType(v.Type))
// 		} else {
// 			tt := []ThingType{}
// 			lx := []Links{}
// 			for _, l := range v.Link {
// 				newLx := Links{
// 					Type:    getLinkType(l.Type),
// 					Value:   l.Value,
// 					ID:      l.ID,
// 					Inbound: l.Inbound,
// 				}
// 				lx = append(lx, newLx)
// 			}
// 			gc[v.ID] = Game{
// 				Title:         v.Name[0].Value,
// 				BggId:         v.ID,
// 				BggType:       append(tt, getThingType(v.Type)),
// 				Yearpublished: v.Yearpublished.Value,
// 				Description:   v.Description,
// 				Thumbnail:     v.Thumbnail,
// 				Image:         v.Image,
// 				Minage:        v.Minage.Value,
// 				Minplayer:     v.Minplayers.Value,
// 				Maxplayer:     v.Maxplayers.Value,
// 				Minplaytime:   v.Minplaytime.Value,
// 				Maxplaytime:   v.Minplaytime.Value,
// 				Links:         lx,
// 			}
// 		}
// 	}

// 	return gc
// }

func (gc *GameCollection) Array() (games []Game) {
	for _, v := range *gc {
		games = append(games, v)
	}

	return games
}
