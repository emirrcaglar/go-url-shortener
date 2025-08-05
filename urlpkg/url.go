package urlpkg

import (
	"database/sql"
	"log"
)

type Url struct {
	ID        int    `json:"id"`
	LONG_URL  string `json:"long_url"`
	SHORT_URL string `json:"short_url"`
}

func (*Url) GenerateShortUrl(db *sql.DB, u *Url, long_url string, baseUrl string, userId int) (string, error) {
	err := saveToDb(db, u, long_url, userId)
	short_url := idToShortUrl(baseUrl, u.ID)
	if err != nil {
		log.Printf("Error saving to DB.")
		return "", err
	}
	u.SHORT_URL = short_url
	err = updateDb(db, u.SHORT_URL, u.LONG_URL, userId)
	if err != nil {
		log.Printf("Error updating DB.")
		return "", err
	}
	return short_url, nil
}
