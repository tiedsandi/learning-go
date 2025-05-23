package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName, lastName, birthdate string
	height                         float64
	createdAt                      time.Time
}

func (u user) outputUserDetails() {
	// fmt.Println(u)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.height)
}

func (u *user) clearUserName(height *float64) {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.height = *height
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname,lastname, and birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
		// height:    userHeight,
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	userHeightStr := getUserData("Please enter your height in cm: ")

	var userHeight float64
	fmt.Sscanf(userHeightStr, "%f", &userHeight)

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// 	// height:    userHeight,
	// }

	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	userHeight,
	// 	time.Now(),
	// }

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)

	appUser.outputUserDetails()
	appUser.clearUserName(&userHeight)
	appUser.outputUserDetails()
}

// func outputUserDetails(u *user) {

// 	fmt.Println(u)
// 	fmt.Println((*u).firstName, u.lastName, u.birthdate)
// }

// func outputUserDetails(firstName, lastName, birthdate string) {
// 	fmt.Println(firstName, lastName, birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
