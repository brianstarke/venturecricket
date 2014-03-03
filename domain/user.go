package domain

import (
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/dancannon/gorethink"
)

var usersTable = "users"

type User struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Wins         uint32 `json:"wins"`
	Losses       uint32 `json:"losses"`
	Abandoned    uint32 `json:"abandoned"`
	TotalMoney   uint32 `json:"totalMoney"`
	PasswordHash string `json:"hashedPassword"`
}

type NewUser struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type UserDomain struct {
	Session *gorethink.Session
}

func (u UserDomain) FindAll() ([]User, error) {
	result := []User{}

	rows, err := gorethink.Table(usersTable).Run(u.Session)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		var u User
		err := rows.Scan(&u)
		if err != nil {
			log.Println(err)
		}
		result = append(result, u)
	}

	return result, nil
}

func (u UserDomain) FindById(id string) (User, error) {
	row, err := gorethink.Table(usersTable).Get(id).RunRow(u.Session)

	if err != nil {
		return User{}, err
	}

	if !row.IsNil() {
		var user User
		row.Scan(&user)

		return user, nil
	} else {
		return User{}, nil
	}
}

func (u UserDomain) CreateUser(newUser *NewUser) (string, error) {
	user := User{}
	user.Username = newUser.Username
	user.EmailAddress = newUser.EmailAddress

	b, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(b)

	resp, err := gorethink.Table(usersTable).Insert(user).RunWrite(u.Session)
	if err != nil {
		return "", err
	} else {
		log.Printf("New user created [%s][%s]", user.Username, user.EmailAddress)
		return resp.GeneratedKeys[0], nil
	}
}
