package main

type ItemMenu struct {
	ID       int
	Kategori string
	Nama     string
	Harga    int
}

type Kategori struct {
	ID   int
	Nama string
}

var kategori = []Kategori{
	{ID: 1, Nama: "Fried Chicken"},
	{ID: 2, Nama: "Geprek"},
	{ID: 3, Nama: "Sadazz"},
	{ID: 4, Nama: "Ayam CLBK"},
	{ID: 5, Nama: "Paket Chicken"},
	{ID: 6, Nama: "Paket Geprek"},
	{ID: 7, Nama: "Side Dish"},
	{ID: 8, Nama: "Minuman"},
}

var daftarMenu = []ItemMenu{

	{Kategori: "Fried Chicken", Nama: "Ayam Dada", Harga: 18000},
	{Kategori: "Fried Chicken", Nama: "Ayam Paha Atas", Harga: 18000},
	{Kategori: "Fried Chicken", Nama: "Ayam Paha Bawah", Harga: 14000},
	{Kategori: "Fried Chicken", Nama: "Ayam Sayap", Harga: 14000},
	{Kategori: "Geprek", Nama: "Ayam Geprek Dada", Harga: 22000},
	{Kategori: "Geprek", Nama: "Ayam Geprek Paha Atas", Harga: 22000},
	{Kategori: "Geprek", Nama: "Ayam Geprek Paha Bawah", Harga: 19500},
	{Kategori: "Geprek", Nama: "Ayam Geprek Sayap", Harga: 18000},
	{Kategori: "Sadazz", Nama: "Sadazz Dada", Harga: 23000},
	{Kategori: "Sadazz", Nama: "Sadazz Paha Atas", Harga: 23000},
	{Kategori: "Sadazz", Nama: "Sadazz Sayap", Harga: 23000},
	{Kategori: "Sadazz", Nama: "Sadazz Paha Bawah", Harga: 23000},
	{Kategori: "Ayam CLBK", Nama: "Ayam CLBK Dada", Harga: 26000},
	{Kategori: "Ayam CLBK", Nama: "Ayam CLBK Paha Atas", Harga: 26000},
	{Kategori: "Ayam CLBK", Nama: "Ayam CLBK Paha Bawah", Harga: 22000},
	{Kategori: "Ayam CLBK", Nama: "Ayam CLBK Sayap", Harga: 22000},
	{Kategori: "Paket Chicken", Nama: "Paket Dada", Harga: 24000},
	{Kategori: "Paket Chicken", Nama: "Paket Paha Atas", Harga: 24000},
	{Kategori: "Paket Chicken", Nama: "Paket Paha Bawah", Harga: 19500},
	{Kategori: "Paket Chicken", Nama: "Paket Sayap", Harga: 19500},
	{Kategori: "Paket Geprek", Nama: "Paket Geprek Paha Atas", Harga: 22000},
	{Kategori: "Paket Geprek", Nama: "Paket Geprek Dada", Harga: 22000},
	{Kategori: "Paket Geprek", Nama: "Paket Geprek Paha Bawah", Harga: 19500},
	{Kategori: "Paket Geprek", Nama: "Paket Geprek Sayap", Harga: 18000},
	{Kategori: "Side Dish", Nama: "Nasi Putih", Harga: 7000},
	{Kategori: "Side Dish", Nama: "Kentang", Harga: 5000},
	{Kategori: "Side Dish", Nama: "Spaghetti", Harga: 5000},
	{Kategori: "Minuman", Nama: "Air Mineral", Harga: 8000},
	{Kategori: "Minuman", Nama: "S-Tee", Harga: 7000},
	{Kategori: "Minuman", Nama: "Teh Botol Kotak Sosro", Harga: 7500},
}
