package main

import (
	"fmt"
)

type barang struct {
	nama  string
	harga int
}

func main() {
	barangs := []barang{
		barang{nama: "GTX 2060 S pki", harga: 5000000},
		barang{nama: "RTX 4090 Ti", harga: 20000000},
		barang{nama: "RX 6950 xt", harga: 15000000},
	}
	//list harga
	fmt.Println("\n Harga")
	for i, barang := range barangs {
		fmt.Printf("%v. %v : Rp %v\n", i+1, barang.nama, barang.harga)
	}

	var transaksi string
	total := 0

	for {
		var pilihBarang, qty, subtotal int

		// pilihan 0 exit
		if pilihBarang == 0 {
			break
		}
		//input barang
		fmt.Print("pilih barang 1 - 3 atau '0' untuk exit : ")
		_, err := fmt.Scanln(&pilihBarang)
		if err != nil {
			fmt.Println("integer only ")
			return
		}
		// memerikasa jika inputan tidak sesuai
		if pilihBarang < 1 || pilihBarang > len(barangs) { //len -> panjang barangs(3)
			fmt.Println("Tolong pilih 1 - 3")
			continue
		}
		// qty
		fmt.Print("Jumlah barang : ")
		_, error := fmt.Scanln(&qty)
		if error != nil {
			fmt.Println("Integer Only")
			return
		}
		// sesuaikan angka pilihan 012 -> 123
		p := barangs[pilihBarang-1]
		subtotal = p.harga * qty

		// penggabungan transaksi
		transaksi += fmt.Sprintf("%v sebanyak %v harganya Rp. %v\n", p.nama, qty, subtotal)

		// menampilkan transaksi
		fmt.Println(transaksi)

		//Total Harga akhir
		total += subtotal
	}
	//menampilkan total
	transaksi += fmt.Sprintf("Total : Rp %v\n", total)
	println(transaksi)
}
