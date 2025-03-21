package main

import (
	"log"
	"os"
	"url_shortner/internal/config"
	"url_shortner/internal/repository"
	"url_shortner/internal/server"
)

func main() {

	// Load Configs.
	cfg := config.Initialize()
	err := repository.Initialize(cfg)
	if err != nil {
		log.Println("Unable to initialize  the repo")
		os.Exit(-1)
	}
	r, err := server.Initialize(cfg)
	if err != nil {
		log.Println("Unable to start the server")
		os.Exit(-1)
	}

	// Start Server.
	err = server.Run(r)
	if err != nil {
		log.Println("Unable to start the server")
		os.Exit(-1)
	}
}
