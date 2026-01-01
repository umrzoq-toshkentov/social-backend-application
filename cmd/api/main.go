package main

import (
	"log"
	"time"

	"github.com/umrzoq-toshkentov/social/internal/db"
	"github.com/umrzoq-toshkentov/social/internal/env"
	"github.com/umrzoq-toshkentov/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5433/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Printf("Db connection estiblashed")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  *store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
