package users

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"github.com/brianstarke/venturecricket/domain"
)

func ListUsers(userDomain *domain.UserDomain, r render.Render) {
	users, err := userDomain.FindAll()

	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, users)
	}

	return
}

func GetUser(userDomain *domain.UserDomain, params martini.Params, r render.Render) {
	user, err := userDomain.FindById(params["id"])

	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, user)
	}

	return
}

type UserAuthRequest struct {
	Query    string `json:"query"`
	Password string `json:"password"`
}

func AuthenticateUser(userDomain *domain.UserDomain, req *http.Request, r render.Render) {
	var u UserAuthRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&u)

	if err != nil {
		r.JSON(400, map[string]interface{}{"error": err.Error()})
	}

	// TODO validation

	user, err := userDomain.Authenticate(u.Query, u.Password)
	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, user)
	}
}

func CreateUser(userDomain *domain.UserDomain, req *http.Request, r render.Render) {
	var u domain.NewUser
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&u)

	if err != nil {
		r.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	// validate new user TODO there's gotta be a better way
	var errors []string

	if u.Username == "" {
		errors = append(errors, "username is required")
	}
	if u.EmailAddress == "" {
		errors = append(errors, "emailAddress is required")
	}
	if u.Password == "" {
		errors = append(errors, "password is required")
	}

	if errors != nil {
		r.JSON(400, map[string]interface{}{"validationErrors": errors})
		return
	}

	id, err := userDomain.CreateUser(&u)

	if err != nil {
		r.JSON(500, map[string]interface{}{"serverError": err.Error()})
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}
