package auth

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB, username, password string) error {
	var userPasswordHash string
	err := db.QueryRow("SELECT userpass FROM users WHERE username = (?)", username).Scan(&userPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return err
	}
	err = checkPasswordHash(userPasswordHash, password)
	if err != nil {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}
