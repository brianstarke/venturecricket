package users

import (
	"log"
	"github.com/codegangsta/martini-contrib/render"
	rethink "github.com/dancannon/gorethink"
)

func ListUsers(session *rethink.Session, r render.Render) {
	result := []User{}
	rows, err := rethink.Table("users").Run(session)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var u User
		err := rows.Scan(&u)
		if err != nil {
			log.Println(err)
		}
		result = append(result, u)
	}
	r.JSON(200, result)
}
