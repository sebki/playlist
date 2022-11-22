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

	for _, v := range res {
		_, err := database.Database.MutateGame(v)
		if err != nil {
			log.Println("error mutating game: ", v.Title, err)
		}
	}

	if err != nil {
		log.Println(err)
	}
}
