package domain

import (
	"log"
	"math/rand"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var cardsTable = "cards"

type CardSuit string

const (
	Red    CardSuit = "RED"
	Green           = "GREEN"
	Blue            = "BLUE"
	Purple          = "PURPLE"
	Yellow          = "YELLOW"
)

/*
Investment cards have a value of 2 to 10.

Multipliers value is always 1.
*/
type CardType string

const (
	Investment CardType = "INVESTMENT"
	Multiplier          = "MULTIPLIER"
)

/*
Since we don't actually have a 'Deck' per se, this determines
where the Card is in the context of the Game.

Deck: combined with Position this determines where in the Deck
this card is.

In Hand: this card has been dealt to a players hand.

In Play: the player has played the card.

Discard: this card has been discarded.  Combined with Position,
this determines where in the discard pile this card is.
*/
type CardLocation string

const (
	Deck    CardLocation = "DECK"
	Hand                 = "HAND"
	InPlay               = "IN_PLAY"
	Discard              = "DISCARD"
)

/*
GameId defines what game this card belongs to.  By querying on this
id, we can pull together the entire deck and game state.

PlayerId is assigned once the card has been assigned to a players Hand
and maps to the Users collection.
*/
type Card struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	GameId   string        `json:"gameId"`
	PlayerId string        `json:"playerId"`
	Suit     CardSuit      `json:"suit"`
	Location CardLocation  `json:"location"`
	Position int           `json:"position"`
	Type     CardType      `json:"type"`
	Value    int           `json:"value"`
}

type CardDomain struct {
	Database *mgo.Database
}

func (d CardDomain) FindAllByGameId(id string) *[]Card {
	c := d.Database.C(cardsTable)

	results := []Card{}

	err := c.Find(bson.M{"gameid": id}).Sort("position").All(&results)

	if err != nil {
		log.Printf("mgo error %s, err is : %s", id, err.Error())
		return nil
	} else {
		return &results
	}
}

/*
Generates a deck for the given gameId
*/
func (d CardDomain) CreateDeck(gameId string) bool {
	// TODO, check if this game already has a deck
	c := d.Database.C(cardsTable)

	suits := [5]CardSuit{Red, Green, Blue, Purple, Yellow}

	// create a randomized slice of deck positions,
	// such that we create a "shuffled" deck
	rand.Seed(time.Now().UTC().UnixNano())
	p := rand.Perm(55)
	for k, _ := range p {
		p[k]++
	}

	for i := 0; i < 5; i++ {
		var suit = suits[i]

		for j := 2; j < 10; j++ {
			x := p[len(p)-1]
			p = p[:len(p)-1]

			card := Card{}
			card.GameId = gameId
			card.Location = Deck
			card.Suit = suit
			card.Type = Investment
			card.Value = j
			card.Position = x

			err := c.Insert(&card)
			if err != nil {
				log.Print(err.Error())
			}
		}

		for l := 0; l < 3; l++ {
			x := p[len(p)-1]
			p = p[:len(p)-1]

			card := Card{}
			card.GameId = gameId
			card.Location = Deck
			card.Suit = suit
			card.Type = Multiplier
			card.Value = 1
			card.Position = x

			err := c.Insert(&card)
			if err != nil {
				log.Print(err.Error())
			}
		}
	}

	return true
}
