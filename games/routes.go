package games

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"

	"github.com/brianstarke/venturecricket/domain"
)

type CreateGameForm struct {
	Player1Id string
}

func CreateGame(gameDomain *domain.GameDomain, req *http.Request, r render.Render) {
	var c CreateGameForm
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&c)

	if err != nil {
		r.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	// validate new user TODO there's gotta be a better way
	var errors []string

	if c.Player1Id == "" {
		errors = append(errors, "player1id is required")
	}

	if errors != nil {
		r.JSON(400, map[string]interface{}{"validationErrors": errors})
		return
	}

	id, err := gameDomain.CreateGame(c.Player1Id)

	if err != nil {
		r.JSON(500, map[string]interface{}{"serverError": err.Error()})
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}
