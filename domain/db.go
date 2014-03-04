package domain

import (
	"log"
	"time"

	"github.com/codegangsta/martini"
	"github.com/dancannon/gorethink"
)

var dbName = "venturecricket"
var dbAddress = "localhost:28015"

func initDB() *gorethink.Session {
	s, err := gorethink.Connect(map[string]interface{}{
		"address":     dbAddress,
		"database":    dbName,
		"maxIdle":     10,
		"idleTimeout": time.Second * 10,
	})

	if err != nil {
		log.Println(err)
	}
	err = gorethink.DbCreate(dbName).Exec(s)
	if err != nil {
		log.Println(err)
	}

	_, err = gorethink.Db(dbName).TableCreate("users").RunWrite(s)
	if err != nil {
		log.Println(err)
	}

	_, err = gorethink.Db(dbName).Table("users").IndexCreate("Username").Run(s)
	if err != nil {
		log.Println(err)
	}

	_, err = gorethink.Db(dbName).Table("users").IndexCreate("EmailAddress").Run(s)
	if err != nil {
		log.Println(err)
	}

	return s
}

/*
Make domain structs available to handlers via middleware
*/
func DomainMiddleware() martini.Handler {
	session := initDB()
	userDomain := &UserDomain{session}

	return func(c martini.Context) {
		c.Map(userDomain)
		c.Next()
	}
}
