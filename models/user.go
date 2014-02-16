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

func UserCreate(id int64, username string, displayName string, avatarUrl string) {
	user := User{
		CreatedAt:   time.Now().UnixNano(),
		Id:          id,
		Username:    username,
		DisplayName: displayName,
		AvatarUrl:   avatarUrl,
	}

	err := dbmap.Insert(&user)
	checkErr(err, "UserCreate failed")
}

func UserFindById(id int64) User {
	user := User{}

	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=?", id)
	checkErr(err, "UserFindById failed")

	return user
}
