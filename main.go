package main

import (
	"os"

	"github.com/farisraii/go-restAPI-echo/db"
	"github.com/farisraii/go-restAPI-echo/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(getPort()))
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
