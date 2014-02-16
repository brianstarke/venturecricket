package models

import (
	"log"
	"time"
)

type User struct {
	Id             int64
	CreatedAt      int64
	Username       string
	Email          string
	HashedPassword string
	Salt           string
	Score          int64
}

func (u *User) Save() {
	u.CreatedAt = time.Now().UnixNano()

	err := dbmap.Insert(u)
	checkErr(err, "User.create failed")
}

func (u *User) FindById(id int64) User {
	user := User{}

	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=?", id)
	checkErr(err, "UserFindById failed")

	return user
}

// initialize the Users table
func init() {
	u := dbmap.AddTableWithName(User{}, "users")
	u.SetKeys(true, "Id")
	u.ColMap("Username").SetUnique(true).SetNotNull(true)
	u.ColMap("Email").SetUnique(true).SetNotNull(true)

	err := dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	log.Println("Users table created")
}
