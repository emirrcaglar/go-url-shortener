package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Couldn't load .env file (using system env vars instead)")
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = os.Getenv("DB_HOST")
	cfg.Addr = os.Getenv("DB_ADDR")
	cfg.DBName = os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println(pingErr)
		return nil, pingErr
	}

	log.Println("Successfully connected to db.")
	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}
