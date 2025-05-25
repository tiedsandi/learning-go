package latihanpointer

import "fmt"

func main() {
	// Latihan 1
	nums := []int{1, 2, 3}
	changeFirstTo100(nums)
	fmt.Println(nums) // [1 2 3]
	changeFirstTo100Ptr(&nums)
	fmt.Println(nums) // [100 2 3]

	// Latihan 2
	user := User{"Joko", 30}
	user.SetNameByValue("Andi")
	fmt.Println(user.Name) // "Joko"
	user.SetNameByPointer("Andi")
	fmt.Println(user.Name) // "Andi"

	// Latihan 3
	products := []*Product{
		{Name: "Pensil"},
		{Name: "Buku"},
	}
	changeProductName(products[0], "Pulpen")
	fmt.Println(products[0].Name) // "Pulpen"

	replaceProducts(products)
	fmt.Println(products[0].Name) // Tetap "Pulpen", tidak berubah

	replaceProductsPtr(&products)
	fmt.Println(products[0].Name) // "Baru1"
}

// 🧤 Latihan 1: Ubah elemen slice
// Tujuan: Pahami efek passing-by-value vs passing-by-pointer ke fungsi.

// ✅ Soal:
// Buat dua fungsi:

// changeFirstTo100(slice []int) → yang tidak mengubah slice asli

// changeFirstTo100Ptr(slice *[]int) → yang mengubah slice asli

// Contoh pemanggilan:

// go
// Salin
// Edit
// nums := []int{1, 2, 3}
// changeFirstTo100(nums)
// fmt.Println(nums) // Tetap [1 2 3]

// changeFirstTo100Ptr(&nums)
// fmt.Println(nums) // Jadi [100 2 3]
// 🏗️ Latihan 2: Struct + Receiver (pointer vs value)
// Tujuan: Pahami perbedaan receiver pointer dan value saat method mengubah struct.

// ✅ Soal:
// Buat struct User:

// go
// Salin
// Edit
// type User struct {
// 	Name string
// 	Age  int
// }
// Buat dua method:

// SetNameByValue(newName string) → pakai receiver value

// SetNameByPointer(newName string) → pakai receiver pointer

// Contoh:

// go
// Salin
// Edit
// user := User{"Joko", 30}
// user.SetNameByValue("Andi")
// fmt.Println(user.Name) // Tetap "Joko"

// user.SetNameByPointer("Andi")
// fmt.Println(user.Name) // Jadi "Andi"
// 🧪 Latihan 3: Slice of pointer vs pointer to slice
// Tujuan: Lihat efek jika kita memodifikasi isi pointer dalam slice vs slice-nya sendiri.

// ✅ Soal:
// Buat slice of pointer ke struct:

// go
// Salin
// Edit
// type Product struct {
// 	Name string
// }

// products := []*Product{
// 	{Name: "Pensil"},
// 	{Name: "Buku"},
// }
// Buat fungsi changeProductName(p *Product, newName string) → ubah nama 1 produk

// Ubah nama elemen pertama jadi "Pulpen" lewat fungsi

// Tantangan ekstra:

// Apa yang terjadi kalau kamu coba ubah slice itu langsung (products = ...) di fungsi?

// Apa bedanya kalau slice-nya dikirim pakai pointer (*[]*Product)?

// Latihan 1
// Contoh penggunaan (bisa dicoba di func main)

func changeFirstTo100(slice []int) {
	if len(slice) > 0 {
		slice[0] = 100 // Ini hanya mengubah copy, tidak mempengaruhi slice asli
	}
}

func changeFirstTo100Ptr(slice *[]int) {
	if len(*slice) > 0 {
		(*slice)[0] = 100 // Ini mengubah slice asli
	}
}

// Latihan 2

type User struct {
	Name string
	Age  int
}

func (u User) SetNameByValue(newName string) {
	u.Name = newName // Hanya mengubah copy, tidak mempengaruhi struct asli
}

func (u *User) SetNameByPointer(newName string) {
	u.Name = newName // Mengubah struct asli
}

// Latihan 3

type Product struct {
	Name string
}

func changeProductName(p *Product, newName string) {
	p.Name = newName
}

// Fungsi untuk mencoba mengubah slice products secara langsung
func replaceProducts(products []*Product) {
	products = []*Product{
		{Name: "Baru1"},
		{Name: "Baru2"},
	}
	// Ini hanya mengubah copy lokal, tidak mempengaruhi slice asli di luar fungsi
}

// Fungsi untuk mengubah slice products lewat pointer
func replaceProductsPtr(products *[]*Product) {
	*products = []*Product{
		{Name: "Baru1"},
		{Name: "Baru2"},
	}
	// Ini akan mengubah slice asli di luar fungsi
}
