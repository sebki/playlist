package bgg

import (
	"log"
	"time"

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

func Crawl() {
	date := time.Now()
	date.Add(-time.Hour * 168)
	links, err := database.Database.GetFamilyLinks(date)
	if err != nil {
		log.Println(err)
	}
	length := len(links)

	for length > 0 {
		log.Println("Found", length, "new links")
		for _, v := range links {
			fq := NewFamilyQuery(v.BggId)
			games, err := Query(fq)
			log.Println("Found", len(games), "Games in Link:", v.LinkValue)
			if err != nil {
				log.Println(err)
			}
			for _, g := range games {
				database.Database.MutateGame(g)
			}
			v.LastBggQuery = time.Now()
			database.Database.MutateLink(v)
		}
		links, err := database.Database.GetFamilyLinks(date)
		if err != nil {
			log.Println(err)
		}
		length = len(links)
	}

	log.Println("Crawl finished")
}
