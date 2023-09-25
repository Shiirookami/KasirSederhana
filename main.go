package main

import (
	"fmt"
	"os"
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
	// list harga
	fmt.Println("\n Harga")
	for i, barang := range barangs {
		fmt.Printf("%v. %v : Rp %v\n", i+1, barang.nama, barang.harga)
	}

	var transaksi string
	total := 0
	pilihBarang := -1

	for pilihBarang != 0 {
		var pilihBarang, qty, subtotal int

		//input barang
		fmt.Print("pilih barang 1 - 3 atau '0' untuk exit : ")
		_, err := fmt.Scanln(&pilihBarang)
		if err != nil {
			fmt.Println("Harap masukkan angka!!")
			return
		}
		// memerikasa jika inputan tidak sesuai
		if pilihBarang < 0 || (pilihBarang != 0 && (pilihBarang < 1 || pilihBarang > len(barangs))) { //len -> panjang barangs(3)
			fmt.Println("Tolong pilih 1 - 3")
			continue
		}

		// pilihan 0 exit
		if pilihBarang == 0 {
			break
		}

		// qty
		fmt.Print("Jumlah barang : ")
		_, error := fmt.Scanln(&qty)
		if error != nil {
			fmt.Println("Harap masukkan angka!!")
			return
		}
		// sesuaikan angka pilihan 012 -> 123
		p := barangs[pilihBarang-1]
		subtotal = p.harga * qty

		// penggabungan transaksi
		transaksi += fmt.Sprintf("%v sebanyak %v harganya Rp. %v\n", p.nama, qty, subtotal)

		//Total Harga akhir
		total += subtotal
	}
	//menampilkan total
	transaksi += fmt.Sprintf("Total : Rp %v\n", total)

	file, err := os.Create("transac.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, writeErr := file.WriteString(transaksi)
	if writeErr != nil {
		fmt.Println("Error:", writeErr)
		return
	}

	fmt.Println("Transaksi telah ditulis ke transac.txt")
	println(transaksi)
}
