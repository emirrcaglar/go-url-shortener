package urlpkg

import (
	"database/sql"
)

type Url struct {
	ID        int    `json:"id"`
	LONG_URL  string `json:"long_url"`
	SHORT_URL string `json:"short_url"`
}

func (*Url) GenerateShortUrl(db *sql.DB, u *Url, long_url, baseUrl string, userID int) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	res, err := tx.Exec("INSERT INTO urls (long_url, short_url, userID) VALUES (?, ?, ?)", long_url, "", userID)
	if err != nil {
		return "", err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	u.ID = int(lid)
	short_url := idToShortUrl(baseUrl, u.ID)
	u.LONG_URL = long_url
	u.SHORT_URL = short_url

	_, err = tx.Exec("UPDATE urls SET short_url = ? WHERE long_url = ? AND userID = ?", short_url, long_url, userID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	// return short_url to print it
	return short_url, nil
}
