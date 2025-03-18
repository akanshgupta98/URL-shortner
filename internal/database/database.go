package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"url_shortner/internal/config"

	_ "github.com/lib/pq"
)

func Initialize(cfg config.Config) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DBCfg.IP, cfg.DBCfg.Port, cfg.DBCfg.User, cfg.DBCfg.Pwd, cfg.DBCfg.DB, cfg.DBCfg.SSLMode)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Unable to initialize the database: ", err.Error())
		return
	}
	log.Println("SUcess to open postgres")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Println("Unable to connect with database: ", err.Error())
		return
	}
	log.Println("Successfully connected to the database")
	return

}
