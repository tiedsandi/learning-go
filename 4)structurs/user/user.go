package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName, LastName, Birthdate string
	Height                         float64
	CreatedAt                      time.Time
}

type Admin struct {
	email, password string
	User
}

func (u User) OutputUserDetails() {
	// fmt.Println(u)
	fmt.Println(u.FirstName, u.LastName, u.Birthdate, u.Height)
}

func (u *User) ClearUserName(Height *float64) {
	u.FirstName = ""
	u.LastName = ""
	u.Birthdate = ""
	u.Height = *Height
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			Birthdate: "---",
			CreatedAt: time.Now(),
			Height:    172,
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("FirstName,LastName, and Birthdate are required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
		// Height:    userHeight,
	}, nil
}
