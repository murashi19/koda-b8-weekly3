package services

import (
	"fmt"
	"strconv"
	"time"
	"weekly3/data"
	"weekly3/utils"
)

func Checkout() {

	if len(data.ShopCart.Items) == 0 {
		fmt.Println("Cart is empty.")
		utils.EnterBack()
		return
	}

	fmt.Println("==========================================")
	fmt.Printf("%28s\n", "ORDER SUMMARY")
	fmt.Println("==========================================")
	fmt.Println()

	fmt.Println("------------------------------------------")

	totalQty := 0
	totalPrice := 0
	for _, item := range data.ShopCart.Items {
		subtotal := item.Quantity * item.Menu.Price

		fmt.Printf(
			"%-22s x%-2d %10s\n",
			item.Menu.Name,
			item.Quantity,
			formatRupiah(subtotal),
		)
		totalQty += item.Quantity
		totalPrice += subtotal
	}

	fmt.Println("------------------------------------------")

	fmt.Printf("%-22s : %d\n", "Items", totalQty)
	fmt.Printf("%-22s : %s\n", "Subtotal", formatRupiah(totalPrice))
	fmt.Printf("%-22s : %s\n", "Tax", formatRupiah(0))
	fmt.Printf("%-22s : %s\n", "Discount", formatRupiah(0))

	fmt.Println("------------------------------------------")

	fmt.Printf("%-22s : %s\n", "TOTAL", formatRupiah(totalPrice))

	fmt.Println("==========================================")

	for {

		payment, err := strconv.Atoi(utils.Input("\nPayment : Rp "))
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		if payment < totalPrice {
			fmt.Println("Payment is insufficient.")
			continue
		}

		change := payment - totalPrice

		done := make(chan struct{})
		go utils.Loading(done, "Payment is being processed...")

		time.Sleep(2 * time.Second)
		close(done)

		fmt.Println("\n    ")
		fmt.Println("==========================================")
		fmt.Println("✔ Payment Successful!")
		fmt.Printf("%-22s : %s\n", "Payment", formatRupiah(payment))
		fmt.Printf("%-22s : %s\n", "Change", formatRupiah(change))
		fmt.Println("==========================================")

		errors := OrderHistory(data.ShopCart.Items, totalQty, totalPrice, payment, change)

		if errors != nil {
			fmt.Println("Failed to save order history:", err)
			utils.EnterBack()
			return
		}
		// Kosongkan cart
		data.ShopCart.Items = nil

		time.Sleep(2 * time.Second)
		return
	}
}
