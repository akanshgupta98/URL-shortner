package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"url_shortner/internal/config"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

var (
	ErrNotInitialized = errors.New("database not initialized")
	ErrInvalidKey     = errors.New("no key exists")
	ErrGeneralFailure = errors.New("unable to fetch information from database")
)

func Initialize(cfg config.ServerConfig) {
	options := &redis.Options{
		Addr: cfg.DB_IP + ":" + cfg.DB_Port,
	}
	rdb = redis.NewClient(options)

}

func Store(key, val string) (err error) {
	if rdb == nil {
		return ErrNotInitialized
	}
	exists := rdb.Exists(context.Background(), key)
	if exists.Val() == 1 {

		log.Println("Key already exists")
		return fmt.Errorf("key: %s already exists", key)

	}
	err = rdb.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		log.Println("Redis store failed: ", err.Error())
		return fmt.Errorf("unable to store in redis db: %s", err.Error())
	}

	return nil

}

func Get(key string) (val string, err error) {
	if rdb == nil {
		return val, ErrNotInitialized
	}
	val, err = rdb.Get(context.Background(), key).Result()
	if err != nil {
		// log.Println("key does not exists")
		if errors.Is(err, redis.Nil) {
			return val, fmt.Errorf("%w: %v", ErrInvalidKey, err)
		} else {
			return val, fmt.Errorf("%w: %v", ErrGeneralFailure, err)
		}

	}
	return val, nil

}
