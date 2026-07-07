package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func Exit() {
	os.Exit(0)
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func EnterBack() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to back...")
	reader.ReadString('\n')
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nError:", r)
			EnterBack()
		}
	}()

	for {
		ClearScreen()

		fmt.Println("========== MENU LAZATTO ==========")

		for _, k := range kategori {
			fmt.Printf("%d. %s\n", k.ID, k.Nama)
		}

		fmt.Println("0. Exit")
		fmt.Println()

		input := Input("Choose category: ")

		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("\nInput harus berupa angka!")
			EnterBack()
			continue
		}

		if id == 0 {
			Exit()
		}

		ShowMenuByCategory(id)
	}
}
