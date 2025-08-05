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
	short_url := idToShortUrl(baseUrl, userId)
	err := saveToDb(db, u, long_url, short_url, userId)
	if err != nil {
		log.Printf("Error saving to DB.")
		return "", err
	}
	return short_url, nil
}
