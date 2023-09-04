package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

func Init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("Couldn't connect to database")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN INVALID")
	}

}

func CreateCon() *sql.DB {
	return db
}
