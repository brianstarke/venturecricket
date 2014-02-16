package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap *gorp.DbMap

func Close() {
	dbmap.Db.Close()
	log.Println("Database connection closed")
}

func init() {
	db, err := sql.Open("mysql", "vcricket:password@/venturecricket")
	checkErr(err, "sql.Open failed")

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.TraceOn("[mysql]", log.New(os.Stdout, "", log.LstdFlags))

	log.Println("Connected to database")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)
	}
}
