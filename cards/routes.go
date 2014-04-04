package cards

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"github.com/brianstarke/venturecricket/domain"
)

func GetCards(cardDomain *domain.CardDomain, params martini.Params, r render.Render) {
	cards, err := cardDomain.FindAllByGameId(params["gameId"])

	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, cards)
	}

	return
}

func CreateDeck(cardDomain *domain.CardDomain, params martini.Params, r render.Render) {
	err := cardDomain.CreateDeck(params["gameId"])

	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, "ok")
	}

	return
}
