package models

type CartItem struct {
	Menu     Menu
	Quantity int
}

type Cart struct {
	Items []CartItem
}
