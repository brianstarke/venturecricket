package domain

import (
	"fmt"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/sendgrid/sendgrid-go"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var usersTable = "users"

type User struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username     string        `json:"username"`
	EmailAddress string        `json:"emailAddress"`
	Wins         uint32        `json:"wins"`
	Losses       uint32        `json:"losses"`
	Abandoned    uint32        `json:"abandoned"`
	TotalMoney   uint32        `json:"totalMoney"`
	PasswordHash string        `json:"passwordHash"`
}

type NewUser struct {
	Username     string
	EmailAddress string
	Password     string
}

type UserDomain struct {
	Database *mgo.Database
}

func (u UserDomain) FindAll() (*[]User, error) {
	c := u.Database.C(usersTable)

	result := []User{}

	err := c.Find(bson.M{}).All(&result)

	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func (u UserDomain) FindById(id string) (*User, error) {
	c := u.Database.C(usersTable)

	result := User{}

	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func (u UserDomain) FindByUsername(query string) *User {
	c := u.Database.C(usersTable)

	result := User{}

	err := c.Find(bson.M{"Username": query}).One(&result)
	log.Print(err)

	if err != nil {
		log.Printf("mgo error looking up %s, err is : %s", query, err.Error())
		return nil
	} else {
		return &result
	}
}

func (u UserDomain) FindByEmailAddress(query string) *User {
	c := u.Database.C(usersTable)

	result := User{}

	err := c.Find(bson.M{"EmailAddress": query}).One(&result)

	if err != nil {
		log.Printf("mgo error looking up %s, err is : %s", query, err.Error())
		return nil
	} else {
		return &result
	}
}

/*
Takes username / email address and attempts to authenticate
against the provided password.  Returns a reference to the User
on success
*/
func (u UserDomain) Authenticate(query string, password string) (*User, error) {
	user := u.FindByEmailAddress(query)

	if user == nil {
		user = u.FindByUsername(query)
	}

	// if we still don't have a the user, bail
	if user == nil {
		return nil, fmt.Errorf("No user found with username/emailAddress %s", query)
	}

	// attempt to authenticate password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("Invalid password for %s", query)
	} else {
		return user, nil
	}
}

func (ud UserDomain) CreateUser(newUser *NewUser) (bool, error) {
	u := ud.FindByUsername(newUser.Username)

	if u != nil {
		return false, fmt.Errorf("User with username %s already exists", u.Username)
	}

	u = ud.FindByEmailAddress(newUser.EmailAddress)

	if u != nil {
		return false, fmt.Errorf("User with emailAddress %s already exists", u.EmailAddress)
	}

	user := User{}
	user.Username = newUser.Username
	user.EmailAddress = newUser.EmailAddress

	b, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return false, err
	}
	user.PasswordHash = string(b)

	c := ud.Database.C(usersTable)

	err = c.Insert(&user)

	if err != nil {
		return false, err
	} else {
		return true, err
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
