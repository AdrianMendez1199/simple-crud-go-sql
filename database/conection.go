package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type connection struct{}

var (
	con  *connection
	once sync.Once
)

func GetInstance() *connection {
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("SINGLETON")
		con = &connection{}
	})

	return con
}

// This function return db connection
func (c *connection) GetConnection() *sql.DB {

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
