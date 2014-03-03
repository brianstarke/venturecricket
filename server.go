package main

import (
	"log"
	"time"

	"github.com/brianstarke/venturecricket/domain"
	"github.com/brianstarke/venturecricket/users"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	rethink "github.com/dancannon/gorethink"
)

var dbName = "venturecricket"
var dbAddress = "localhost:28015"

// TODO move all this to domain package
func InitDB() *rethink.Session {
	s, err := rethink.Connect(map[string]interface{}{
		"address":     dbAddress,
		"database":    dbName,
		"maxIdle":     10,
		"idleTimeout": time.Second * 10,
	})

	if err != nil {
		log.Println(err)
	}
	err = rethink.DbCreate(dbName).Exec(s)
	if err != nil {
		log.Println(err)
	}

	_, err = rethink.Db(dbName).TableCreate("users").RunWrite(s)
	if err != nil {
		log.Println(err)
	}

	return s
}

func DB() martini.Handler {
	session := InitDB()
	userDomain := &domain.UserDomain{session}

	return func(c martini.Context) {
		c.Map(userDomain)
		c.Map(session)
		c.Next()
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/", func() string {
		return "WORK IN PROGRESS"
	})

	// user routes
	m.Get("/users", users.ListUsers)
	m.Get("/users/:id", users.GetUser)
	m.Post("/users", users.CreateUser)

	m.Run()
}
