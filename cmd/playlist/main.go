package main

import (
	"github.com/sebki/playlist/internal/database"
	"github.com/sebki/playlist/internal/server"
)

func main() {
	db := database.NewClient()
	defer db.Closer()

	database.Database = db
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

	server.Start(":3030")

}
