package bgg

import (
	"log"

	"github.com/sebki/playlist/internal/database"
)

func AddHotness() {
	hq := NewHotQuery("boardgame")
	res, err := Query(hq)
	if err != nil {
		log.Println(err)
	}

	err = database.Database.CreateGames(res...)

	if err != nil {
		log.Println(err)
	}
}
