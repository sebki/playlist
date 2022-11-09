package server

import "github.com/gin-gonic/gin"

func Start(port string) {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/bggsearch", bggsearch)
	r.GET("/bgghotness", bgghotness)

	r.Run(port)
}
