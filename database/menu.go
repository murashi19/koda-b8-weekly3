package database

import "weekly3/models"

var Categories = []models.Category{
	{ID: 1, Name: "Fried Chicken"},
	{ID: 2, Name: "Burger"},
	{ID: 3, Name: "Geprek"},
	{ID: 4, Name: "Combo"},
	{ID: 5, Name: "Side Dish"},
	{ID: 6, Name: "Beverage"},
}

var ShopCart models.Cart

var Menus = []models.Menu{

	// Fried Chicken
	{ID: 1, CategoryID: 1, Name: "Original Sayap", Price: 12000, Stock: 50},
	{ID: 2, CategoryID: 1, Name: "Original Paha Bawah", Price: 14000, Stock: 50},
	{ID: 3, CategoryID: 1, Name: "Original Paha Atas", Price: 16000, Stock: 50},
	{ID: 4, CategoryID: 1, Name: "Original Dada", Price: 16000, Stock: 50},

	// Burger
	{ID: 5, CategoryID: 2, Name: "Burger Reguler", Price: 11000, Stock: 40},
	{ID: 6, CategoryID: 2, Name: "Burger Cheese", Price: 13000, Stock: 40},
	{ID: 7, CategoryID: 2, Name: "Burger Premium Cheese", Price: 17000, Stock: 30},
	{ID: 8, CategoryID: 2, Name: "Burger Double Beef", Price: 15000, Stock: 30},

	// Geprek
	{ID: 9, CategoryID: 3, Name: "Geprek Sayap", Price: 18000, Stock: 30},
	{ID: 10, CategoryID: 3, Name: "Geprek Paha Bawah", Price: 19000, Stock: 30},
	{ID: 11, CategoryID: 3, Name: "Geprek Paha Atas", Price: 21000, Stock: 30},
	{ID: 12, CategoryID: 3, Name: "Geprek Dada", Price: 21000, Stock: 30},

	// Combo
	{ID: 13, CategoryID: 4, Name: "Combo Original 1", Price: 19000, Stock: 25},
	{ID: 14, CategoryID: 4, Name: "Combo Original 2", Price: 21000, Stock: 25},
	{ID: 15, CategoryID: 4, Name: "Combo CLBK", Price: 23000, Stock: 20},
	{ID: 16, CategoryID: 4, Name: "Sadazz 1", Price: 23000, Stock: 20},

	// Side Dish
	{ID: 17, CategoryID: 5, Name: "French Fries", Price: 9500, Stock: 40},
	{ID: 18, CategoryID: 5, Name: "Spaghetti", Price: 13000, Stock: 30},
	{ID: 19, CategoryID: 5, Name: "Hot Dog", Price: 10500, Stock: 30},
	{ID: 20, CategoryID: 5, Name: "Nasi", Price: 6000, Stock: 100},

	// Beverage
	{ID: 21, CategoryID: 6, Name: "S-Tee", Price: 4000, Stock: 100},
	{ID: 22, CategoryID: 6, Name: "Lemon Tea", Price: 6500, Stock: 80},
	{ID: 23, CategoryID: 6, Name: "Air Mineral", Price: 4000, Stock: 100},
	{ID: 24, CategoryID: 6, Name: "Milo", Price: 6000, Stock: 50},
}
