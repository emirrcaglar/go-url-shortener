package auth

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func checkUserName(db *sql.DB, username string) bool {
	_, err := db.Exec("SELECT * FROM users WHERE username = (?)", username)
	return err != nil
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func checkPasswordHash(pw, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	return err
}
