package db

import (
	"database/sql"
	"log"

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
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connected to database!")

	return DB, err
}
