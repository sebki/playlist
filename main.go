package main

import (
	"fmt"
	"log"
)

func main() {
	c, close := newClient()
	defer close()

	u, err := createNewUser(c, "sebki4000", "masmas@torn-relic.de", "!Basti4Online")
	if err != nil {
		if !IsValidationError(err) {
			log.Fatal(err)
		}
		fmt.Println(err)
	}
	fmt.Println(u)

	u, err = loginByEmail(c, "info@sleeved.de", "!Basti4Online")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()
}
