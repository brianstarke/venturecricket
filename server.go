package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/brianstarke/venturecricket/users"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	rethink "github.com/dancannon/gorethink"
)

var dbName = "venturecricket"
var dbAddress = "localhost:28015"

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

	return func(c martini.Context) {
		c.Map(session)
		c.Next()
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/", func() string {
		return "Hey there."
	})

	// user routes
	m.Get("/users", users.ListUsers)
	m.Post("/users", HandleNewUser)

	m.Run()
}

// test
type NewUser struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
}

func HandleNewUser(req *http.Request, r render.Render) {
	decoder := json.NewDecoder(req.Body)
	var u NewUser
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	log.Print(u)

	r.JSON(200, map[string]interface{}{"user added": "ok"})
}
