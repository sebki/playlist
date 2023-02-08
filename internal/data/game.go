package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sebki/playlist/internal/validator"
)

type Game struct {
	ID            int64           `json:"id"`
	CreatedAt     time.Time       `json:"-"`
	ModifiedAt    time.Time       `json:"-"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	YearPublished int32           `json:"year_published"`
	Type          []BoardgameType `json:"type"`
	Thumbnail     string          `json:"thumbnail,omitempty"`
	Image         string          `json:"image,omitempty"`
	MinPlayer     int32           `json:"min_player"`
	MaxPlayer     int32           `json:"max_plaxer"`
	MinPlaytime   int32           `json:"min_playtime"`
	MaxPlaytime   int32           `json:"max_playtime"`
	MinAge        int32           `json:"min_age"`
	MaxAge        int32           `json:"max_age"`
	// Tags          *[]Tag          `json:"tags,omitempty"`
	// Persons       *[]Person       `json:"persons"`
	// Orgs          *[]Organization `json:"organizations"`
	Version int32 `json:"version"`
}

type BoardgameType string

const (
	Standalone BoardgameType = "standalone"
	Expansion  BoardgameType = "expansion"
)

func (bgt BoardgameType) String() string {
	switch bgt {
	case Standalone:
		return string(bgt)
	case Expansion:
		return string(bgt)
	default:
		return "not valid"
	}
}

func checkTypes(bgt []BoardgameType) error {
	for _, v := range bgt {
		if v.String() == "not valid" {
			return errors.New("not a valid BoardgameType")
		}
	}

	return nil
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(game.Description != "", "description", "must be provided")
	v.Check(game.YearPublished >= 0, "year_published", "must be greater than 0")
	v.Check(game.YearPublished <= int32(time.Now().Year()+1), "year_published", "maximum one year ahead allowed")
	v.Check(checkTypes(game.Type) != nil, "type", "not a valid type")
	v.Check(game.MinPlayer >= 1, "min_player", "must not be less than 1")
	v.Check(game.MinPlayer <= 200, "min_player", "must be less than 200")
	v.Check(game.MaxPlayer >= 1, "max_player", "must not be less than 1")
	v.Check(game.MaxPlayer <= 200, "max_player", "must be less than 200")
	v.Check(game.MinPlaytime >= 0, "min_playtime", "must not be less than 0")
	v.Check(game.MinPlaytime <= 1000, "min_playtime", "must be less than 1000")
	v.Check(game.MaxPlaytime >= 0, "max_playtime", "must not be less than 0")
	v.Check(game.MaxPlaytime < 1000, "max_playtime", "must be less than 1000")
	v.Check(game.MinAge >= 0, "min_age", "must not be less than 0")
	v.Check(game.MinAge < 150, "min_age", "must be less than 150")
	v.Check(game.MaxAge >= 0, "max_age", "must not be less than 0")
	v.Check(game.MaxAge < 150, "max_age", "must be less than 150")
}

type GameModel struct {
	DB *sql.DB
}
