package main

import (
	"fmt"
	"log"
)

func main() {
	c, close := newClient()
	defer close()

	u := User{
		Username: "sebki3000",
		Email:    "info@sleeved.de",
		Password: "!Basti4Online",
	}

	err := u.createNewUser(c)
	if err != nil {
		if !IsValidationError(err) {
			log.Fatal(err)
		}
		fmt.Println(err)
	}

	err = u.login(c)
	if err != nil {
		log.Fatal(err)
	}

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()
}
