package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// This function return db connection
func GetConnection() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	dns := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("BD_NAME") + "?sslmode=disable"

	db, err := sql.Open("postgres", dns)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
