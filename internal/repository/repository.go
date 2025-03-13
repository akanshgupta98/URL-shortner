package repository

import (
	"context"
	"fmt"
	"log"
	"url_shortner/internal/config"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Initialize(cfg config.ServerConfig) {
	options := &redis.Options{
		Addr: cfg.DB_IP + ":" + cfg.DB_Port,
	}
	rdb = redis.NewClient(options)

}

func Store(key, val string) (err error) {

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
	val, err = rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("key does not exists")
		return val, fmt.Errorf("no url exists with this key: %s", key)
	}
	return val, nil

}
