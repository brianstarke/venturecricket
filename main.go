package main

import (
	"log"

	"github.com/brianstarke/venturecricket/models"
)

func main() {
	defer models.Close()

	models.InitializeTables()

	u1 := models.User{
		Id:          1,
		Username:    "brianstarke",
		DisplayName: "Brian Starke",
		AvatarUrl:   "http://www.google.com",
	}

	u1.Save()

	user := models.User{}
	user = user.FindById(1)
	log.Println(user.DisplayName)
}
