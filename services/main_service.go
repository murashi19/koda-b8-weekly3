package services

import (
	"fmt"
	"strconv"
	"weekly3/data"
	"weekly3/models"
	"weekly3/utils"
)

func ShowMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("===========================================")
		fmt.Println("      🍗 LAZATTO CHICKEN & BURGER")
		fmt.Println("===========================================")
		fmt.Println("")

		fmt.Println("Welcome to Lazatto Ordering System")
		fmt.Println("")

		fmt.Println("1. Order Food")
		fmt.Println("2. Search Menu")
		fmt.Println("3. View Cart")
		fmt.Println("4. Checkout")
		fmt.Println("0. Exit")
		fmt.Println("")

		input, err := strconv.Atoi(utils.Input("Choose: "))
		if err != nil {
			fmt.Println("Failed to convert Number")
			utils.EnterBack()
			continue
		}
		switch input {
		case 1:
			ShowCategory()
		case 3:
			utils.ClearScreen()
			ViewCart()
		case 0:
			return

		default:
			fmt.Println("Menu not found.")
			utils.EnterBack()
		}
	}

}

func ShowCategory() {
	for {
		utils.ClearScreen()

		fmt.Println("====== CATEGORY ======")

		for _, c := range data.Categories {
			fmt.Printf("%d. %s\n", c.ID, c.Name)
		}

		fmt.Println("0. Back")

		input, err := strconv.Atoi(utils.Input("Choose: "))
		if err != nil {
			fmt.Println("Invalid input!")
			utils.EnterBack()
			continue
		}

		if input == 0 {
			return
		}

		ShowItem(input)
	}
}

func ShowItem(categoryID int) {
	utils.ClearScreen()

	var categoryName string

	for _, c := range data.Categories {
		if c.ID == categoryID {
			categoryName = c.Name
			break
		}
	}

	if categoryName == "" {
		fmt.Println("Category not found.")
		utils.EnterBack()
		return
	}

	fmt.Println("======", categoryName, "======")

	for _, menu := range data.Menus {
		if menu.CategoryID == categoryID {
			fmt.Printf("%d. %s - Rp%d\n",
				menu.ID,
				menu.Name,
				menu.Price,
			)
		}
	}

	fmt.Println("0. Back")

	input, err := strconv.Atoi(utils.Input("Choose Menu ID: "))
	if err != nil {
		fmt.Println("Invalid input!")
		utils.EnterBack()
		return
	}

	if input == 0 {
		return
	}

	// addCart
	AddToCart(categoryID, input)
}

func FindMenuByID(id int) models.Menu {
	for _, menu := range data.Menus {
		if menu.ID == id {
			return menu
		}
	}

	return models.Menu{}
}
