package domain

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var gamesTable = "games"

type GameState string

const (
	Pending    GameState = "PENDING"
	InProgress           = "IN_PROGRESS"
	Complete             = "COMPLETE"
)

/*
GameId defines what game this card belongs to.  By querying on this
id, we can pull together the entire deck and game state.

PlayerId is assigned once the card has been assigned to a players Hand
and maps to the Users collection.
*/
type Game struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Player1Id string        `json:"player1Id"`
	Player2Id string        `json:"player2Id" bson:"omitempty"`
	State     GameState     `json:"state"`
}

type GameDomain struct {
	Database *mgo.Database
}

func (d GameDomain) CreateGame(player1Id string) (*string, error) {
	c := d.Database.C(gamesTable)

	game := Game{}
	game.Id = bson.NewObjectId()
	game.Player1Id = player1Id
	game.State = Pending

	gameId := game.Id.Hex()
	err := c.Insert(&game)

	if err != nil {
		return nil, err
	} else {
		return &gameId, nil
	}
}
