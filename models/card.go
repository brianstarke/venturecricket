package models

import (
	"log"

	"github.com/coopernurse/gorp"
)

type Card struct {
	Id        int64
	CreatedAt int64
	Position  int32  // the cards position in the deck/discard pile
	Location  string // one of DECK/DISCARD/P1/P2
}

// initialize the Cards table
func (card *Card) Init(dbmap *gorp.DbMap) {
	c := dbmap.AddTableWithName(Card{}, "cards")
	c.SetKeys(true, "Id")

	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	log.Println("Cards table created")
}
