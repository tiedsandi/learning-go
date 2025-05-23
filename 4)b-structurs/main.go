package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

type fullName struct {
	firstName, lastName string
}

func (f fullName) Introducing(age int) {
	fmt.Println("My name is", f.firstName, f.lastName, "and this year my age is", age)
}

func main() {
	var name customString = "sandi"

	name.log()

	var myFullname fullName

	myFullname = fullName{
		firstName: "Fachran",
		lastName:  "Sandi",
	}

	myFullname.Introducing(25)
}
