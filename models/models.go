package models

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var dbmap *gorp.DbMap

// TODO move these definitions to their own files
type User struct {
	Id          int64 // the users twitter ID
	CreatedAt   int64
	Username    string
	DisplayName string
	AvatarUrl   string
}

// creates all the tables needed, unless they already exist
func InitializeTables() {
	dbmap.AddTableWithName(User{}, "users")

	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
}

func Close() {
	dbmap.Db.Close()
	log.Println("Database connection closed")
}

func init() {
	db, err := sql.Open("mysql", "vcricket:password@/vcricket")
	checkErr(err, "sql.Open failed")

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	log.Println("Connected to database")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)
	}
}
