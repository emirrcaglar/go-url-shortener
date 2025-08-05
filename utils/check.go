package utils

import (
	"fmt"

	"github.com/emirrcaglar/go-url-shortener/session"
)

func CheckStatus() (*session.Cfg, error) {
	cfg, err := session.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load session: %w", err)
	}

	if !cfg.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	return cfg, nil
}
