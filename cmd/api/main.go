package main

import (
	"log"

	"github.com/codeboris/specialist/internal/app/api"
)

var ()

func init() {}

func main() {
	log.Println("It`s works!")

	server := api.New()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
