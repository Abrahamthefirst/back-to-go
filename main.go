package main

import (
	"github.com/Abrahamthefirst/back-to-go/application"
)

func main() {
	// db := database.NewPgDB()
	app := application.NewApp()
	app.Bootstrap()

}
