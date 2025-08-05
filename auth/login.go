package auth

import (
	"database/sql"
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/types"
)

func Login(db *sql.DB, username, password string) (*types.User, error) {
	var userPasswordHash string
	var id int

	err := db.QueryRow("SELECT id, userpass FROM users WHERE username = (?)", username).Scan(&id, &userPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid username or password")
		}
		return nil, err
	}
	err = checkPasswordHash(userPasswordHash, password)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	return &types.User{
		ID:       id,
		UserName: username,
	}, nil
}
