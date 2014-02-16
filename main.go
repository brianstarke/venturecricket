package main

import "github.com/brianstarke/venturecricket/models"

func main() {
	defer models.Close()

	models.InitializeTables()
}
