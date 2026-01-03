package main

import (
	"fmt"
	"log"
	"time"

	"github.com/umrzoq-toshkentov/social/internal/db"
	"github.com/umrzoq-toshkentov/social/internal/env"
	"github.com/umrzoq-toshkentov/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5433/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, time.Minute*15)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
	fmt.Println("Completed seed")
}
