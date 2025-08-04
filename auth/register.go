package auth

import (
	"database/sql"
	"fmt"
)

func Register(db *sql.DB, username, password string) error {
	if checkUserName(db, username) {
		fmt.Errorf("username already exists")
	}

	pw, err := hashPassword(password)
	if err != nil {
		fmt.Errorf("error hashing password", err)
	}

	_, err = db.Exec("INSERT INTO users (username, userpass) VALUES (?, ?)", username, string(pw))
	if err != nil {
		fmt.Errorf("error registering user", err)
	}
	return nil
}
