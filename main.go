package main

import (
	"github.com/ardin2001/go_mini-capstone/routes"
	"github.com/ardin2001/go_mini-capstone/utils"
)

func main() {
	utils.UserMigrate()
	e := routes.StartApp()
	e.Logger.Fatal(e.Start(":8000"))
}
