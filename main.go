package main

import (
	"log"

	"github.com/brianstarke/venturecricket/models"
)

func main() {
	defer models.Close()

	u1 := models.User{
		Username: "brianstarke",
		Email:    "brian.starke@gmail.com",
	}

	u1.Save()

	user := models.User{}
	user = user.FindById(1)
	log.Println(user.Username)
}
