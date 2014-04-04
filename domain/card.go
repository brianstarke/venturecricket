package domain

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var cardsTable = "cards"

type CardSuit int

const (
	Red CardSuit = iota
	Green
	Blue
	Purple
	Yellow
)

type Card struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	GameId bson.ObjectId `json:"gameId"`
	Suit   CardSuit      `json:"suit"`
}

type CardDomain struct {
	Database *mgo.Database
}
