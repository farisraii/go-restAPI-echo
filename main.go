package main

import (
	"github.com/farisraii/go-restAPI-echo/db"
	"github.com/farisraii/go-restAPI-echo/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
