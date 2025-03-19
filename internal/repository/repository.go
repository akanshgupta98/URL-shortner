package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"url_shortner/internal/config"
	"url_shortner/internal/database"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

var (
	ErrNotInitialized = errors.New("database not initialized")
	ErrBadKey         = errors.New("no key exists")
	ErrFetch          = errors.New("unable to fetch information from database")
	ErrDuplicateKey   = errors.New("key already exists in the database")
	ErrStore          = errors.New("unable to store in the db")
	ErrClose          = errors.New("unable to close the db connection")
)

type Repo struct {
	DBHdlr *sql.DB
	once   sync.Once
}

var r Repo

func Initialize(cfg config.Config) (err error) {

	r.once.Do(func() {
		r.DBHdlr, err = database.Initialize(cfg)
		if err != nil {
			log.Println("Unable to initialize repo: ", err.Error())

		}
		log.Println("Repository initialized")
	})
	if err != nil {
		return fmt.Errorf("%w %v", ErrNotInitialized, err)
	}

	return

}

func Store(key, val string) error {
	if r.DBHdlr == nil {
		log.Println(ErrNotInitialized.Error())
		return ErrNotInitialized
	}

	query := fmt.Sprintf(`INSERT INTO SHORTNER (SHORT_URL,OG_URL) VALUES ('%s','%s')`, key, val)
	rows, err := r.DBHdlr.Query(query)
	if err != nil {
		log.Println("Database store failed: ", err.Error())
		return fmt.Errorf("%w: %v", ErrStore, err)
	}
	defer rows.Close()

	return nil

}

func Get(key string) (val string, err error) {
	if rdb == nil {
		return val, ErrNotInitialized
	}
	val, err = rdb.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return val, fmt.Errorf("%w: %v", ErrBadKey, err)
		} else {
			return val, fmt.Errorf("%w: %v", ErrFetch, err)
		}

	}
	return val, nil

}
