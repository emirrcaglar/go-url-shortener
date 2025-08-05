package urlpkg

import (
	"database/sql"
	"fmt"
)

type Url struct {
	ID        int    `json:"id"`
	LONG_URL  string `json:"long_url"`
	SHORT_URL string `json:"short_url"`
}

func (*Url) GenerateShortUrl(db *sql.DB, u *Url, long_url, baseUrl string, userID int) (string, error) {

	short_url, err := u.checkExistingUrl(db, long_url, userID)
	if err != nil {
		return "", err
	}
	if short_url != "" {
		return short_url, nil
	}

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
	short_url = idToShortUrl(baseUrl, u.ID)
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

func (*Url) checkExistingUrl(db *sql.DB, long_url string, userID int) (string, error) {
	row := db.QueryRow("SELECT long_url, short_url, userID FROM urls WHERE long_url = ? AND userID = ?", long_url, userID)
	var (
		dblongUrl, shortCode string
		dbUserID             int
	)
	err := row.Scan(&dblongUrl, &shortCode, &dbUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return shortCode, nil
}

func GenerateCustomUrl(db *sql.DB, u *Url, long_url, custom, baseUrl string, userID int) error {
	// Check if the custom short code already exists
	exists, err := u.checkExistingCustomUrl(db, custom, userID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("custom URL '%s' already exists", custom)
	}

	_, err = db.Exec("INSERT INTO urls (long_url, short_url, userID) VALUES (?, ?, ?)", long_url, custom, userID)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Generated custom URL: %s%s\n", baseUrl, custom)
	return nil
}

// New function to check if custom short URL already exists
func (*Url) checkExistingCustomUrl(db *sql.DB, short_url string, userID int) (bool, error) {
	row := db.QueryRow("SELECT short_url FROM urls WHERE short_url = ? AND userID = ?", short_url, userID)
	var dbShortUrl string
	err := row.Scan(&dbShortUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
