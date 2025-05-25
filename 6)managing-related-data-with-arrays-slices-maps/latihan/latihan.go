package main

import (
	"fmt"
)

func main() {

	// 1
	hobbies := []string{"Reading", "Gaming", "Walking"}

	fmt.Println("My Hobbies :", hobbies)
	// 2
	firstHobby := hobbies[0:1]
	fmt.Println("First Hobbie :", firstHobby)

	newHobbies := hobbies[1:3]
	fmt.Println("My new Hobbies :", newHobbies)

	// 3
	fmt.Println(len(firstHobby), cap(firstHobby))
	mainHobbies := hobbies[0:2]
	usingFirstHobbies := firstHobby[0:2]

	fmt.Println(mainHobbies, usingFirstHobbies)

	// 4
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	var courseGoal []string = []string{"Mastering Go-lang", "Can creat rest-api"}

	fmt.Println(courseGoal)

	// 6
	courseGoal[1] = "Make rest-api base on golang"
	fmt.Println(courseGoal)

	courseGoal = append(courseGoal, "Can help to get a job")
	fmt.Println(courseGoal)

	// 7
	type product struct {
		id    int
		title string
		price float64
	}

	products := []product{{id: 1, title: "Buku", price: 3000}, {id: 2, title: "Pensil", price: 1500}}

	fmt.Println(products)

	products = append(products, product{id: 3, title: "Penghapus", price: 500})
	fmt.Println(products)

	products2 := []product{
		{
			1,
			"A First Product",
			12.99,
		},
		{
			2,
			"A Second Product",
			129.99,
		},
	}

	fmt.Println(products2)

	newProduct := product{
		3,
		"A Third Product",
		15.99,
	}

	products2 = append(products2, newProduct)

	fmt.Println(products2)
}

// func main2() {
// 	// 1) Array dengan tiga hobi
// 	var hobbies [3]string = [3]string{"Coding", "Cycling", "Photography"}
// 	fmt.Println("My Hobbies:", hobbies)

// 	// 2) Elemen pertama, dan elemen kedua & ketiga sebagai slice
// 	fmt.Println("First Hobby:", hobbies[0])
// 	secondAndThird := hobbies[1:3]
// 	fmt.Println("Second and Third Hobbies:", secondAndThird)

// 	// 3) Slice berdasarkan elemen pertama yang berisi elemen pertama & kedua (dua cara)
// 	slice1 := hobbies[0:2]
// 	slice2 := hobbies[:2]
// 	fmt.Println("Slice1:", slice1)
// 	fmt.Println("Slice2:", slice2)

// 	// 4) Slice ulang untuk berisi elemen kedua dan terakhir
// 	slice3 := hobbies[1:3]
// 	fmt.Println("Slice3 (second and last):", slice3)

// 	// 5) Dynamic array (slice) tujuan kursus
// 	courseGoals := []string{"Understand Go basics", "Build web API"}
// 	fmt.Println("Course Goals:", courseGoals)

// 	// 6) Ubah tujuan kedua dan tambah tujuan ketiga
// 	courseGoals[1] = "Create REST API with Go"
// 	courseGoals = append(courseGoals, "Get a Go developer job")
// 	fmt.Println("Updated Course Goals:", courseGoals)

// 	// 7) Struct Product dan dynamic list
// 	type Product struct {
// 		id    int
// 		title string
// 		price int
// 	}
// 	products := []Product{
// 		{id: 1, title: "Laptop", price: 10000000},
// 		{id: 2, title: "Mouse", price: 150000},
// 	}
// 	fmt.Println("Products:", products)
// 	products = append(products, Product{id: 3, title: "Keyboard", price: 300000})
// 	fmt.Println("Updated Products:", products)
// }

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// Saatnya latihan dari apa yang sudah kamu pelajari!

// 1) Buat array baru (!) yang berisi tiga hobi kamu
//    Tampilkan (print) array tersebut di command line.
// 2) Tampilkan juga data lain dari array tersebut:
//    - Elemen pertama (sendiri)
//    - Elemen kedua dan ketiga digabung sebagai list baru
// 3) Buat slice berdasarkan elemen pertama yang berisi
//    elemen pertama dan kedua.
//    Buat slice tersebut dengan dua cara berbeda (jadi ada dua slice di akhir)
// 4) Slice ulang slice dari (3) dan ubah agar berisi elemen kedua
//    dan terakhir dari array asli.
// 5) Buat "array dinamis" yang berisi tujuan kursus kamu (minimal 2 tujuan)
// 6) Ubah tujuan kedua menjadi tujuan yang berbeda DAN tambahkan tujuan ketiga ke array dinamis yang sudah ada
// 7) Bonus: Buat struct "Produk" dengan title, id, price dan buat
//    list dinamis produk (minimal 2 produk).
//    Lalu tambahkan produk ketiga ke list produk yang sudah ada.
