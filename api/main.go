package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"api/internal/config"
	"api/internal/handlers"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

var (
	host = "localhost"
	port = 5432
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := connectInProd(config)
	if err != nil {
		log.Fatal(err)
	}

	yapHandler := handlers.NewYapHandler(db)

	r := chi.NewRouter()

	// Mount routes
	r.Put("/trades/{id}/yap", yapHandler.IncreaseYapCounter)

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func connectInDev(config *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword, host, port, config.PostgresDB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, err
}

func connectInProd(config *config.Config) (*sql.DB, error) {
	prodHost := "db"

	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword, prodHost, port, config.PostgresDB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatal(err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatal(err)
	}
	return db, err
}
