package main

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// c, close := newClient()
	// defer close()

	// u, err := createNewUser(c, "sebki4000", "masmas@torn-relic.de", "!Basti4Online")
	// if err != nil {
	// 	if !IsValidationError(err) {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(err)
	// }
	// fmt.Println(u)

	// u, err = loginByEmail(c, "info@sleeved.de", "!Basti4Online")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(u)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/bggsearch", bggsearch)

	r.Run(":3030")
}
