package service

import (
	"errors"
	"fmt"
	"log"
	"url_shortner/internal/repository"

	"github.com/google/uuid"
)

var ErrInvalidRequest = errors.New("invalid data")
var ErrFailure = errors.New("failed to fetch data")

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
	if err != nil {
		if errors.Is(err, repository.ErrBadKey) {
			return originalURL, fmt.Errorf("%w: %v", ErrInvalidRequest, err)
		} else {
			return originalURL, fmt.Errorf("%w: %v", ErrFailure, err)
		}
	}
	return

}
