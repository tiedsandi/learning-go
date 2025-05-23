package lists

import "fmt"

func main() {
	// jadi initnya array di go itu dia overwrite, data lama akan ditiban dengan data baru
	prices := []float64{10.99, 90.22}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	// cara menambahkan array baru
	// prices[2] = 11.99
	updatedPrices := append(prices, 5.88)
	prices = append(prices, 5.88)

	fmt.Println("tambah array baru:", updatedPrices, prices)

	// cara menghapus nilai di array
	prices = prices[1:]

	fmt.Println("hapus nilai di array:", updatedPrices, prices)

	// new section
	updatedWithMutiplePrices := append(prices, 5.88, 5.22, 123.2, 512.341)
	fmt.Println("tambah banyak data ke array:", updatedWithMutiplePrices)

	discountPrices := []float64{101.99, 20.3, 33.3, 44.6}
	newPricesArr := append(prices, discountPrices...)
	fmt.Println(" data ke array:", newPricesArr)

}

// func main() {
// 	var productName [4]string = [4]string{"A book"}

// 	prices := [4]float64{10.88, 9.99, 34.52, 20.2}

// 	fmt.Println("prices:", prices)
// 	fmt.Println("prices[1]:", prices[1])

// 	productName[2] = "A carpet"
// 	fmt.Println("productName:", productName)

// 	featuredPrices := prices[1:3] // slice mengambil elemen dari indeks 1 sampai sebelum indeks 3 (yaitu indeks 1 dan 2)
// 	featuredPrices1 := prices[:3] // slice mengambil elemen dari awal sampai sebelum indeks 3 (yaitu indeks 0, 1, dan 2)
// 	featuredPrices2 := prices[1:] // slice mengambil elemen dari indeks 1 sampai akhir (yaitu indeks 1, 2, dan 3)
// 	featuredPrices2[0] = 199.99   // akan menimpa nilai index 1 pada array prices karena slice featuredPrices2 mereferensikan array yang sama
// 	fmt.Println("featuredPrices (prices[1:3]):", featuredPrices)
// 	fmt.Println("featuredPrices1 (prices[:3]):", featuredPrices1)
// 	fmt.Println("featuredPrices2 (prices[1:]):", featuredPrices2)

// 	fmt.Println("prices setelah diubah lewat featuredPrices2:", prices)

// 	fmt.Println(len(featuredPrices2), cap(featuredPrices2))
// 	// len = 3 (karena ada 3 elemen: indeks 1, 2, 3)
// 	// cap = 3 (karena slice dimulai dari indeks 1 sampai akhir array prices)

// 	highlightedPrices := featuredPrices2[:1] // highlightedPrices adalah slice yang mengambil elemen pertama dari featuredPrices2 (yaitu elemen pada indeks 1 dari prices)
// 	fmt.Println("highlightedPrices (featuredPrices2[:1]):", highlightedPrices)

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// 	// len = 1 (hanya 1 elemen di slice)
// 	// cap = 3 (kapasitas tetap dihitung dari posisi awal featuredPrices2 ke akhir array prices)

// 	// Gunakan len untuk mengetahui berapa banyak data yang bisa diakses di slice.
// 	// Gunakan cap untuk memahami seberapa banyak slice bisa diperluas sebelum perlu membuat array baru di belakang layar.

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println("highlightedPrices (featuredPrices2[:1]):", highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
