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
	repository.Initialize()

	// Start Server.
	err := server.Run(cfg)
	if err != nil {
		log.Println("Unable to start the server")
		os.Exit(-1)
	}
}
