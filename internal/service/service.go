package service

import (
	"fmt"
	"log"
	"url_shortner/internal/repository"

	"github.com/google/uuid"
)

func URLShortner(inputData string) (shortURL string, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("failed to generate UUID: %v\n", err)
		return
	}

	shortURL = id.String()[:7]
	log.Println("URL shortened is: ", shortURL)
	err = repository.Store(shortURL, inputData)
	return shortURL, err

}

func URLShortnerFetch(inputData string) (originalURL string, err error) {
	originalURL, err = repository.Get(inputData)
	return originalURL, err

}
