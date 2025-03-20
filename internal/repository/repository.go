package repository

import (
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

// Strictly for testing purpose only.
func resetDB() {
	r.once = sync.Once{}
	r.DBHdlr = nil

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
	log.Println("KEY IS: ", key)
	query := fmt.Sprintf(`SELECT OG_URL FROM SHORTNER WHERE SHORT_URL='%s'`, key)
	rows, err := r.DBHdlr.Query(query)
	if err != nil {
		log.Println("Not able to fetch: ", err.Error())
		return val, fmt.Errorf("%w: %v", ErrFetch, err)
	}
	defer rows.Close()
	var url string
	for rows.Next() {
		log.Println("INSIDE FOR")
		err = rows.Scan(&url)
		if err != nil {
			return val, fmt.Errorf("%w: %v", ErrFetch, err)
		}
		log.Println("Data is:  ", url)
		val = url
	}
	if val == "" {
		return val, fmt.Errorf("%w", ErrBadKey)
	}
	return

}
