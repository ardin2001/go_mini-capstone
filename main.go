package main

import (
	"github.com/ardin2001/go_mini-capstone/routes"
)

func main() {
	e := routes.StartApp()
	e.Logger.Fatal(e.Start(":8000"))
}
