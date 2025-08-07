package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(HOST, PORT, USER, PASSWORD, DBNAME, SSLMODE string) (*sql.DB, error) {
	connStr := "host=" + HOST + " port=" + PORT + " user=" + USER + " password=" + PASSWORD + " dbname=" + DBNAME + " sslmode=" + SSLMODE
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	// Tentativas de conexão com backoff
	for i := 0; i < 10; i++ {
		err = DB.Ping()
		if err == nil {
			break
		}
		log.Printf("Waiting for database to be ready (attempt %d)...", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	log.Println("Connected to database!")

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration", // Caminho relativo para a pasta das migrations
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations ran successfully!")

	return DB, err
}
