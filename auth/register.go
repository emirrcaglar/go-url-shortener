package auth

import (
	"database/sql"
	"fmt"
)

func Register(db *sql.DB, username, password string) error {
	if checkUserName(db, username) {
		fmt.Println("invalid username")
	}

	pw, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, userpass) VALUES (?, ?)", username, string(pw))
	if err != nil {
		return err
	}
	return nil
}
