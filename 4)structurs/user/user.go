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
