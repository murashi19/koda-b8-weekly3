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
			done := make(chan struct{})

			go utils.Loading(done, "Adding to cart...")
			item.Quantity += qty
			fmt.Printf(
				"✔ %s Quantity updated to %d\n",
				menu.Name,
				item.Quantity,
			)
			time.Sleep(1 * time.Second)
			close(done)
			return
		}
	}
	done := make(chan struct{})

	go utils.Loading(done, "Adding to cart...")
	data.ShopCart.Items = append(
		data.ShopCart.Items,
		models.CartItem{
			Menu:     menu,
			Quantity: qty,
		},
	)
	fmt.Println("===================")
	fmt.Println("✔", menu.Name, "Rp", menu.Price, "x", qty)

	time.Sleep(2 * time.Second)
	close(done)

	fmt.Println("\nSuccessfully added to the cart")
	time.Sleep(1 * time.Second)
}

func ViewCart() {

	for {
		utils.ClearScreen()
		fmt.Println("=======================================================")
		fmt.Printf("%26s\n", "YOUR CART")
		fmt.Println("=======================================================")
		fmt.Println()

		if len(data.ShopCart.Items) == 0 {
			fmt.Println("Your cart is empty")
			utils.EnterBack()
			return
		}

		fmt.Printf("%-4s %-24s %-8s %-5s %-10s\n",
			"No",
			"Menu",
			"Menu ID",
			"Qty",
			"Price",
		)
		fmt.Println("-------------------------------------------------------")

		totalQty := 0
		totalPrice := 0
		for i, cart := range data.ShopCart.Items {
			subtotal := cart.Quantity * cart.Menu.Price
			fmt.Printf(
				"%-4d %-24s %-8d %-5d %-10s\n",
				i+1,
				cart.Menu.Name,
				cart.Menu.ID,
				cart.Quantity,
				formatRupiah(cart.Menu.Price),
			)

			fmt.Printf(
				"%-24s %-12s %-6s %-10s\n",
				"Subtotal",
				"",
				"",
				formatRupiah(subtotal),
			)

			fmt.Println()

			totalQty += cart.Quantity
			totalPrice += subtotal
		}
		fmt.Println("\n-------------------------------------------------------")
		fmt.Printf("Total Item : %d\n", totalQty)
		fmt.Printf("Total Price : %s\n", formatRupiah(totalPrice))
		fmt.Println("")

		fmt.Println("=======================================================")
		fmt.Println("")

		fmt.Println("1. Update Quantity")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Checkout")
		fmt.Println("0. Back")

		choose, err := strconv.Atoi(utils.Input("Choose : "))
		if err != nil {
			fmt.Println("Invalid input!")
			utils.EnterBack()
			continue
		}
		switch choose {
		case 1:
			UpdateItem()
		case 2:

		case 3:
			utils.ClearScreen()
			Checkout()
		case 0:
			return
		default:
			fmt.Println("Invalid input!")
		}
		utils.EnterBack()
	}

}

func UpdateItem() {
	for {
		choose, err := strconv.Atoi(utils.Input("Choose Menu ID: "))
		if err != nil {
			fmt.Println("Please input a valid number.")
			continue
		}

		index := FindCartItemIndex(choose)
		if index == -1 {
			fmt.Println("Menu not found in cart.")
			continue
		}

		item := &data.ShopCart.Items[index]

		qty, err := strconv.Atoi(utils.Input("Update Quantity to: "))
		if err != nil {
			fmt.Println("Please input a valid number.")
			continue
		}

		if qty <= 0 {
			fmt.Println("Quantity must be at least 1.")
			continue
		}

		if qty > item.Menu.Stock {
			fmt.Println("Stock is not enough.")
			continue
		}

		item.Quantity = qty

		fmt.Printf("✔ %s quantity updated to %d\n",
			item.Menu.Name,
			item.Quantity,
		)

		time.Sleep(time.Second)
		return
	}
}

func RemoveItem() {
	for {
		choose, err := strconv.Atoi(utils.Input("Choose Menu ID: "))
		if err != nil {
			fmt.Println("Please input a valid number.")
			continue
		}

		index := FindCartItemIndex(choose)
		if index == -1 {
			fmt.Println("Menu not found in cart.")
			continue
		}

		item := data.ShopCart.Items[index]

		data.ShopCart.Items = append(
			data.ShopCart.Items[:index],
			data.ShopCart.Items[index+1:]...,
		)

		fmt.Printf("✔ %s removed from cart.\n", item.Menu.Name)

		time.Sleep(time.Second)

		return
	}
}

func FindCartItemIndex(menuID int) int {
	for i := range data.ShopCart.Items {
		if data.ShopCart.Items[i].Menu.ID == menuID {
			return i
		}
	}
	return -1
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
