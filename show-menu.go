package main

import (
	"fmt"
)

func ShowMenuByCategory(id int) {
	ClearScreen()

	var namaKategori string

	for _, k := range kategori {
		if k.ID == id {
			namaKategori = k.Nama
			break
		}
	}

	if namaKategori == "" {
		fmt.Println("Kategori tidak ditemukan.")
		EnterBack()
		return
	}

	fmt.Println("======", namaKategori, "======")

	no := 1
	for _, menu := range daftarMenu {
		if menu.Kategori == namaKategori {
			fmt.Printf("%d. %s - Rp%d\n",
				no,
				menu.Nama,
				menu.Harga,
			)
			no++
		}

	}

	fmt.Println()
	EnterBack()
}
