package main

import (
	"fmt"
)

type list struct {
	Nama  string
	Harga float64
	Stock int
}

func main() {

	listBarang := []list{
		{"Roti coklat", 25.0, 70},
		{"Selai Strawberry", 30.0, 80},
		{"Permen", 26.0, 60},
		{"Ice cream", 27.0, 75},
	}

	// Menampilkan daftar barang
	fmt.Println("Daftar barang yang tersedia:")
	for i, list := range listBarang {
		fmt.Printf("%d. %s (Harga: %.2f, Stock: %d)\n", i+1, list.Nama, list.Harga, list.Stock)
	}

	// Inisialisasi variabel total dan keranjang belanja
	var total float64
	keranjangBelanja := make(map[string]int)

	// Inputan pengguna
	for {
		var pilihan int
		var jumlah int

		fmt.Print("Pilih barang (nomor) atau ketik 0 untuk selesai belanja: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			break
		}

		if pilihan < 1 || pilihan > len(listBarang) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		fmt.Print("Jumlah barang yang dibeli: ")
		fmt.Scan(&jumlah)

		barang := listBarang[pilihan-1]

		if jumlah > barang.Stock {
			fmt.Println("Stok tidak mencukupi.")
			continue
		}

		// Menambahkan barang ke keranjang belanja
		keranjangBelanja[barang.Nama] += jumlah

		// Mengurangi stok barang
		listBarang[pilihan-1].Stock -= jumlah

		// Menghitung total
		total += barang.Harga * float64(jumlah)
	}

	// Menampilkan total dan cetak hasil total
	fmt.Printf("Total biaya: %.2f\n", total)

	fmt.Println("Rincian Belanja:")
	for barang, jumlah := range keranjangBelanja {
		fmt.Printf("%s: %d\n", barang, jumlah)
	}
}
