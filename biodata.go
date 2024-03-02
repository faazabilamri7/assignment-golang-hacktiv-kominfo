package main

import (
	"fmt"
	"os"
	"strconv"
)

// Teman struct to store friend's data
type Teman struct {
	Nama             string
	Alamat           string
	Pekerjaan        string
	AlasanMemilihGo  string 
}

// DataTeman is a map to store friends' data by their absen number
var DataTeman = map[int]Teman{
	1: {"John Doe", "Jakarta", "Developer", "Saya suka dengan kecepatan dan kesederhanaan Golang."},
	2: {"Jane Doe", "Surabaya", "Designer", "Golang sangat efisien untuk pengembangan aplikasi."},
	// Add more  friends here...
}

// ShowBiodata function to display friend's data by absen number
func ShowBiodata(absen int) {
	if teman, found := DataTeman[absen]; found {
		fmt.Printf("Data teman dengan absen %d:\n", absen)
		fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan Memilih Golang: %s\n", teman.Nama, teman.Alamat, teman.Pekerjaan, teman.AlasanMemilihGo)
	} else {
		fmt.Println("Teman dengan absen tersebut tidak ditemukan.")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go <absen>")
		os.Exit(1)
	}

	absen := os.Args[1]
	fmt.Printf("Mencari data teman dengan absen %s...\n", absen)

	// Handle error when converting string to integer
	if absenInt, err := strconv.Atoi(absen); err == nil {
		ShowBiodata(absenInt)
	} else {
		fmt.Println("Masukkan absen dalam bentuk angka.")
	}
}
