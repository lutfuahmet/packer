package main

import (
	"log"
	"packer/api"
	"packer/config"
)

func main() {
	// read config file
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	api.StartServer(appConfig)
}
