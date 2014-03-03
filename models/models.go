package models

import "log"

func init() {
	log.Println("INIT")
}

type User struct {
	Username       string `json:"username"`
	EmailAddress   string `json:"emailAddress"`
	Wins           uint32 `json:"wins"`
	Losses         uint32 `json:"losses"`
	Abandoned      uint32 `json:"abandoned"`
	TotalMoney     uint32 `json:"totalMoney"`
	HashedPassword string `json:"hashedPassword"`
	Salt           string `json:"salt"`
}
