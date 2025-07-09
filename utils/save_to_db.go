package utils

import (
	"log"

	"github.com/emirrcaglar/go-url-shortener/db"
)

func saveToDb(u *Url, url string) (err error) {
	db, err := db.Connect()
	if err != nil {
		log.Printf("Error connecting to data at saveToDB.")
		return err
	}

	defer db.Close()

	res, err := db.Exec(
		"INSERT INTO url (long_url) VALUES (?)", url)
	if err != nil {
		log.Printf("Error inserting into table. %v\n", err)
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting LastInsertId.")
		return err
	}

	u.ID = int(lid)
	u.LONG_URL = url
	return nil
}
