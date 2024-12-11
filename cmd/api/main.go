package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/codeboris/specialist/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("It`s works!")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs file:", err)
	}
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
