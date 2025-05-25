package main

import "fmt"

// alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // type, default, max array
	// userNames := []string{}

	userNames[0] = "Joko"

	userNames = append(userNames, "Fachran")
	userNames = append(userNames, "sandi")

	fmt.Println(userNames)

	// new section

	// courseRating := make(map[string]float64, 3)
	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.9
	courseRating["node"] = 4.9
	courseRating["laravel"] = 4.9

	courseRating.output()
	// fmt.Println(courseRating)

}
