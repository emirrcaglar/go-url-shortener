package url

import (
	"database/sql"
	"log"
)

func saveToDb(db *sql.DB, u *Url, url string) (err error) {
	res, err := db.Exec(
		"INSERT INTO urls (long_url) VALUES (?)", url)
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
