package cards

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"github.com/brianstarke/venturecricket/domain"
)

func GetCards(cardDomain *domain.CardDomain, params martini.Params, r render.Render) {
	cards := cardDomain.FindAllByGameId(params["gameId"])

	if cards == nil {
		r.JSON(500, "No cards matching that game id")
	} else {
		r.JSON(200, cards)
	}

	return
}

func CreateDeck(cardDomain *domain.CardDomain, params martini.Params, r render.Render) {
	res := cardDomain.CreateDeck(params["gameId"])

	if res {
		r.JSON(200, "WIN")
	} else {
		r.JSON(500, "FAIL")
	}

	return
}
