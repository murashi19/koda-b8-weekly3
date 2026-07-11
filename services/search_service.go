package services

import (
	"fmt"
	"strconv"
	"strings"

	"weekly3/data"
	"weekly3/models"
	"weekly3/utils"
)

func SearchMenu() {

	utils.ClearScreen()

	fmt.Println("==========================================")
	fmt.Println("              SEARCH MENU")
	fmt.Println("==========================================")

	keyword := strings.TrimSpace(
		strings.ToLower(
			utils.Input("Search : "),
		),
	)

	var results []models.Menu

	for _, menu := range data.Menus {
		if strings.Contains(
			strings.ToLower(menu.Name),
			keyword,
		) {
			results = append(results, menu)
		}
	}

	fmt.Println()

	if len(results) == 0 {
		fmt.Println("No menu found.")
		utils.EnterBack()
		return
	}

	fmt.Println("\n==========================================")
	fmt.Println("             SEARCH RESULT")
	fmt.Println("==========================================")
	fmt.Printf("Keyword : %s\n\n", keyword)

	fmt.Printf("%-4s %-25s %-10s\n", "ID", "Menu", "Price")
	fmt.Println("------------------------------------------")

	for _, menu := range results {
		fmt.Printf(
			"%-4d %-25s %-10s\n",
			menu.ID,
			menu.Name,
			formatRupiah(menu.Price),
		)
	}

	fmt.Println("------------------------------------------")

	fmt.Println("0. Back to Menu")
	fmt.Println("Add to cart, then select the menu ID listed above")

	for {
		choose, err := strconv.Atoi(utils.Input("\nChoose ID: "))
		if err != nil {
			fmt.Println("Invalid input, please enter an existing ID or 0")
			continue
		}
		if choose == 0 {
			ShowMenu()
		}
		found := false
		for _, item := range results {
			if choose == item.ID {
				AddToCart(item.CategoryID, item.ID)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Invalid ID choice.")
			continue
		}
	}
}
