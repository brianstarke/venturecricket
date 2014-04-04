package domain

import (
	"github.com/codegangsta/martini"
	"labix.org/v2/mgo"
)

var dbAddress = "localhost"
var dbName = "vc"

func initDB() *mgo.Database {
	session, err := mgo.Dial(dbAddress)

	if err != nil {
		panic(err)
	} else {
		return session.DB(dbName)
	}
}

/*
Make domain structs available to handlers via middleware
*/
func DomainMiddleware() martini.Handler {
	db := initDB()
	userDomain := &UserDomain{db}
	cardDomain := &CardDomain{db}

	return func(c martini.Context) {
		c.Map(userDomain)
		c.Map(cardDomain)
		c.Next()
	}
}
