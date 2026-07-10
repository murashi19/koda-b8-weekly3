package main

import (
	"os"
	"weekly3/services"
)

func Exit() {
	os.Exit(0)
}

func main() {
	services.ShowMenu()
}
