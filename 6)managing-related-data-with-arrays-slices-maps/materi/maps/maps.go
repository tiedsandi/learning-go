package maps

import "fmt"

type Product struct {
	Id    string
	Title string
	Price float64
}

func main() {
	// website := []string{"google.com", "aws.com"}

	// Tidak bisa menggunakan tipe campuran (string | int) sebagai key pada map di Go.
	// Solusinya, gunakan satu tipe data saja, misal string:
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	// menambahkan map baru
	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	// hapus map
	delete(websites, "LinkedIn")
	delete(websites, "Google")
	fmt.Println(websites)

	store := map[int]Product{
		1: {Id: "produk 1", Title: "Beras", Price: 2000},
		2: {Id: "produk 1", Title: "Beras", Price: 2000},
	}

	fmt.Println(store)
}
