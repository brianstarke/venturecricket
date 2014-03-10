package domain

import (
	"fmt"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/dancannon/gorethink"
	"github.com/sendgrid/sendgrid-go"
)

var usersTable = "users"

type User struct {
	Id           string `json:"id" gorethink:"id,omitempty"`
	Username     string `json:"username" gorethink:"username"`
	EmailAddress string `json:"emailAddress" gorethink:"emailAddress"`
	Wins         uint32 `json:"wins" gorethink:"wins"`
	Losses       uint32 `json:"losses" gorethink:"losses"`
	Abandoned    uint32 `json:"abandoned" gorethink:"abandoned"`
	TotalMoney   uint32 `json:"totalMoney" gorethink:"totalMoney"`
	PasswordHash string `json:"hashedPassword" gorethink:"passwordHash"`
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
		var user User
		err := rows.Scan(&user)
		if err != nil {
			log.Println(err)
		}
		result = append(result, user)
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

func (u UserDomain) FindByUsername(query string) (*User, error) {
	rows, err := gorethink.Table(usersTable).GetAllByIndex("Username", query).Run(u.Session)

	if err != nil {
		return nil, err
	}

	if !rows.IsNil() {
		var user User

		rows.Next()
		rows.Scan(&user)

		return &user, nil
	} else {
		return nil, nil
	}
}

func (u UserDomain) FindByEmailAddress(query string) (*User, error) {
	rows, err := gorethink.Table(usersTable).GetAllByIndex("EmailAddress", query).Run(u.Session)

	if err != nil {
		return nil, err
	}

	if !rows.IsNil() {
		var user User

		rows.Next()
		rows.Scan(&user)

		return &user, nil
	} else {
		return nil, nil
	}
}

/*
Takes username / email address and attempts to authenticate
against the provided password.  Returns a reference to the User
on success
*/
func (u UserDomain) Authenticate(query string, password string) (*User, error) {
	user, err := u.FindByEmailAddress(query)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = u.FindByUsername(query)

		if err != nil {
			return nil, nil
		}
	}

	// if we still don't have a the user, bail
	if user == nil {
		return nil, fmt.Errorf("No user found with username/emailAddress %s", query)
	}

	// attempt to authenticate password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("Invalid password for %s", query)
	} else {
		return user, nil
	}
}

func (ud UserDomain) CreateUser(newUser *NewUser) (string, error) {
	u, err := ud.FindByUsername(newUser.Username)
	if err != nil {
		return "", err
	}
	if u != nil {
		return "", fmt.Errorf("User with username %s already exists", u.Username)
	}

	u, err = ud.FindByEmailAddress(newUser.EmailAddress)
	if err != nil {
		return "", err
	}
	if u != nil {
		return "", fmt.Errorf("User with emailAddress %s already exists", u.EmailAddress)
	}

	user := User{}
	user.Username = newUser.Username
	user.EmailAddress = newUser.EmailAddress

	b, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(b)

	resp, err := gorethink.Table(usersTable).Insert(user).RunWrite(ud.Session)
	if err != nil {
		return "", err
	} else {
		log.Printf("New user created [%s][%s]", user.Username, user.EmailAddress)
		log.Print(resp)
		return resp.GeneratedKeys[0], nil
	}
}

// TODO move account info to .env, add more params
func (u UserDomain) SendPasswordResetEmail(resetId string) {
	sendgridUser := "brianstarke"
	sendgridPass := "iO9w5dcQ1SUR"
	sg := sendgrid.NewSendGridClient(sendgridUser, sendgridPass)

	message := sendgrid.NewMail()
	message.AddTo("brian.starke@gmail.com")
	message.AddToName("Brian Starke")
	message.AddSubject("Password reset link for Venturecricket")
	message.AddText(resetId)
	message.AddFrom("venturecricket@dogfort.io")
	message.AddFromName("venturecricket")
	if r := sg.Send(message); r == nil {
		log.Println("Email sent!")
	} else {
		log.Println(r)
	}
}
