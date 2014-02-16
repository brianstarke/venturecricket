package main

import (
	"github.com/brianstarke/vcricket/models"
)

func main() {
	defer models.Close()

	models.InitializeTables()
}
