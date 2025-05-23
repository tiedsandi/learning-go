package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	userHeightStr := getUserData("Please enter your height in cm: ")

	var userHeight float64
	fmt.Sscanf(userHeightStr, "%f", &userHeight)

	// appUser, err := &user.User(userFirstName, userLastName, userBirthdate)

	// var appUser *user.User
	// appUser = &user.User{
	// 	FirstName: userFirstName,
	// }

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

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

	appUser.OutputUserDetails()
	appUser.ClearUserName(&userHeight)
	appUser.OutputUserDetails()

	// appUser.outputUserDetails()
	// appUser.clearUserName(&userHeight)
	// appUser.outputUserDetails()
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
