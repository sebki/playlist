package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sebki/playlist/internal/bgg"
	"github.com/sebki/playlist/internal/errors"
	"github.com/sebki/playlist/internal/models"
)

func bggsearch(c *gin.Context) {
	searchTerm, isExist := c.GetQuery("query")
	if !isExist {
		log.Println("No query found")
		return
	}

	q := bgg.NewSearchQuery(searchTerm)

	tt := []models.ThingType{
		models.TypeBoardGame,
		models.TypeBoardGameExpansion,
	}
	q.SetThingType(tt)

	res, err := bgg.Query(q)
	if err != nil {
		errors.InternalServerError(c, err)
	}

	c.JSON(200, res.Array())
}
