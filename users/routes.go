package users

import (
	"log"

	"github.com/brianstarke/venturecricket/models"
	"github.com/codegangsta/martini-contrib/render"
	rethink "github.com/dancannon/gorethink"
)

func ListUsers(session *rethink.Session, r render.Render) {
	result := []models.User{}
	rows, err := rethink.Table("users").Run(session)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u)
		if err != nil {
			log.Println(err)
		}
		result = append(result, u)
	}
	r.JSON(200, result)
}
