package auth

import (
	"database/sql"
	"fmt"
)

func login(db *sql.DB, username, password string) error {
	var userPasswordHash string
	err := db.QueryRow("SELECT userpass FROM users WHERE username = (?)", username).Scan(userPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return fmt.Errorf("database error: %w", err)
	}
	err = checkPasswordHash(userPasswordHash, password)
	if err != nil {
		return fmt.Errorf("invalid username or password")
	}
	_, err = db.Exec("INSERT INTO users(username, userpass) VALUES (?, ?)", username, password)
	if err != nil {
		return fmt.Errorf("registration error")
	}
	return nil
}
