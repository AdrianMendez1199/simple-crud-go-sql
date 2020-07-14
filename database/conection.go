package database

import (
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type connection struct{}

var (
	con  *connection
	once sync.Once
)

// Singleton
func GetInstance() *connection {
	// 1 intance form stuct connection
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Fatal(err)
		}

		con = &connection{}
	})

	return con
}

// This function return db connection
func (c *connection) GetConnection() *gorm.DB {

	dns := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("BD_NAME") + "?sslmode=disable"

	db, err := gorm.Open("postgres", dns)

	// if app in debug mode show log sql query
	if os.Getenv("GO_ENV") == "development" {
		db.LogMode(true)
	}

	if err != nil {
		panic(err)
	}

	return db
}
