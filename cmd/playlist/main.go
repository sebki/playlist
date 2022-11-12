package main

import (
	"github.com/sebki/playlist/internal/bgg"
	"github.com/sebki/playlist/internal/database"
	"github.com/sebki/playlist/internal/server"
)

func main() {
	db := database.NewClient()
	defer db.Closer()

	database.Database = db

	bgg.AddHotness()

	server.Start(":3030")

}
