package services

import (
	"fmt"
	"strconv"
	"time"
	"weekly3/data"
	"weekly3/models"
	"weekly3/utils"
)

func AddToCart(categoryID int, menuID int) {

	menu := FindMenuByID(menuID)

	if menu.CategoryID != categoryID {
		fmt.Println("Invalid menu for this category.")
		utils.EnterBack()
		return
	}
	fmt.Println("You selected:", menu.Name)
	qty, err := strconv.Atoi(utils.Input("Quantity: "))
	if err != nil {
		fmt.Println("Invalid quantity. Please enter a valid number.")
		utils.EnterBack()
		return
	}
	fmt.Printf("Qty : %d\n", qty)

	if qty <= 0 {
		fmt.Println("Quantity cannot be less than 1.")
		utils.EnterBack()
		return
	}
	if qty > menu.Stock {
		fmt.Println("Stock is not enough.")
		utils.EnterBack()
		return
	}
	for i := range data.ShopCart.Items {
		item := &data.ShopCart.Items[i]
		if item.Menu.ID == menu.ID {
			if item.Quantity+qty > menu.Stock {
				fmt.Println("Requested quantity exceeds available stock.")
				utils.EnterBack()
				return
			}
			item.Quantity += qty
			fmt.Printf(
				"✔ %s Quantity updated to %d\n",
				menu.Name,
				item.Quantity,
			)
			time.Sleep(2 * time.Second)
			return
		}
	}
	data.ShopCart.Items = append(
		data.ShopCart.Items,
		models.CartItem{
			Menu:     menu,
			Quantity: qty,
		},
	)
	fmt.Println("===================")
	fmt.Println("✔", menu.Name, "Rp", menu.Price, "x", qty)
	fmt.Println("Added to Cart")
	time.Sleep(1 * time.Second)
}

func ViewCart() {

	for {
		fmt.Println("==========================================")
		fmt.Printf("%26s\n", "YOUR CART")
		fmt.Println("==========================================")
		fmt.Println()

		if len(data.ShopCart.Items) == 0 {
			fmt.Println("Your cart is empty")
			utils.EnterBack()
			return
		}

		fmt.Printf("%-4s %-20s %5s %10s\n", "No", "Menu", "Qty", "Price")
		fmt.Println("------------------------------------------")

		totalQty := 0
		totalPrice := 0

		item := &data.ShopCart.Items
		for i, cart := range *item {
			subtotal := cart.Quantity * cart.Menu.Price
			fmt.Printf(
				"%-4d %-20s %5d %10s\n",
				i+1,
				cart.Menu.Name,
				cart.Quantity,
				formatRupiah(cart.Menu.Price),
			)

			fmt.Printf(
				"     %-20s %5s %10s\n",
				"Subtotal",
				"",
				formatRupiah(subtotal),
			)

			fmt.Println()

			totalQty += cart.Quantity
			totalPrice += subtotal
		}
		fmt.Println("\n------------------------------------------")
		fmt.Printf("Total Item : %d\n", totalQty)
		fmt.Printf("Total Price : %s\n", formatRupiah(totalPrice))
		fmt.Println("")

		fmt.Println("==========================================")
		fmt.Println("")
		fmt.Println("9. Checkout Now")
		choose, err := strconv.Atoi(utils.Input("Choose : "))
		if err != nil || choose != 9 {
			fmt.Println("Invalid input!")
			utils.EnterBack()
			continue
		}
		utils.EnterBack()
	}

}

func formatRupiah(n int) string {
	s := strconv.Itoa(n)

	var result string
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		result = string(s[i]) + result
		count++

		if count%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp" + result
}
