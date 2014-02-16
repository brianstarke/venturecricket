package models

import (
	"time"
)

type User struct {
	Id          int64 // the users twitter ID
	CreatedAt   int64
	Username    string
	DisplayName string
	AvatarUrl   string
	Score       int32
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
