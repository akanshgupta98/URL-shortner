package repository

import (
	"fmt"
	"log"
	"sync"
)

var dataStore map[string]string
var mu sync.Mutex

func Initialize() {
	dataStore = make(map[string]string)

}

func Store(key, val string) (err error) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := dataStore[key]; ok {

		log.Println("Key already exists")
		return fmt.Errorf("key: %s already exists", key)

	}
	dataStore[key] = val
	return nil

}

func Get(key string) (val string, err error) {
	mu.Lock()
	defer mu.Unlock()
	var ok bool
	if val, ok = dataStore[key]; !ok {
		log.Println("key does not exists")
		return val, fmt.Errorf("no url exists with this key: %s", key)
	}
	return val, nil

}
