package urlpkg

import (
	"database/sql"
	"log"
)

func saveToDb(db *sql.DB, u *Url, long_url string, userId int) (err error) {
	res, err := db.Exec(
		"INSERT INTO urls (long_url, short_url, userID) VALUES (?, ?, ?)", long_url, "", userId)
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
	u.LONG_URL = long_url
	return nil
}
