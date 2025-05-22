package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age
	// agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("*Age Pointer:", *agePointer)
	fmt.Println("Age Pointer:", agePointer)

	fmt.Println("getAdult Years :", getAdultYears(age))
	fmt.Println("getAdult Years Pointers:", getAdultYearsPointer(agePointer))

	editAgeToAdultYears(agePointer)
	fmt.Println("getAdult Years Pointers without print:", age)
}

func getAdultYears(age int) int {
	return age - 19
}
func getAdultYearsPointer(age *int) int {
	return *age - 19
}
func editAgeToAdultYears(age *int) {
	*age = *age - 18
}

// func main() {
//     x := 10         // Ini rumah, isinya 10
//     p := &x         // Ini alamat rumah, pointer ke x

//     fmt.Println("x:", x)      // 10
//     fmt.Println("p:", p)      // alamat di memori, misal 0xc000012080
//     fmt.Println("*p:", *p)    // 10, isi dari alamat tsb

//     *p = 20                  // ubah isi rumah lewat alamat
//     fmt.Println("x:", x)     // 20
// }
