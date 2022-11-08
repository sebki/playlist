package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sebki/playlist/internal/bgg"
	"github.com/sebki/playlist/internal/errors"
)

func bgghotness(c *gin.Context) {
	hq := bgg.NewHotQuery("boardgame")
	res, err := bgg.Query(hq)
	if err != nil {
		errors.InternalServerError(c, err)
	}

	c.JSON(200, res)
}
