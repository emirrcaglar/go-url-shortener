package session

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/emirrcaglar/go-url-shortener/types"
)

type Cfg struct {
	LoggedIn    bool        `json:"loggedIn"`
	CurrentUser *types.User `json:"currentUser"`
	CreatedAt   time.Time   `json:"createdAt"`
	Expiration  time.Time   `json:"expiration"`
}

var ConfigFile string

func init() {
	home, _ := os.UserHomeDir()
	ConfigFile = filepath.Join(home, ".urlshortener", "config.json")
}

func NewSession(user *types.User, ttl time.Duration) *Cfg {
	now := time.Now()
	return &Cfg{
		LoggedIn:    true,
		CurrentUser: user,
		CreatedAt:   now,
		Expiration:  now.Add(ttl),
	}
}

func LoadConfig() (*Cfg, error) {
	cfg := &Cfg{}
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil // return empty config
		}
		return nil, err
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.HasExpired() {
		cfg.LoggedIn = false
		cfg.CurrentUser = nil
	}

	return cfg, nil
}

func (c *Cfg) HasExpired() bool {
	return c.LoggedIn && time.Now().After(c.Expiration)
}

func SaveConfig(cfg *Cfg) error {
	// Ensure dir exists
	_ = os.MkdirAll(filepath.Dir(ConfigFile), 0755)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigFile, data, 0644)
}

func Clear() error {
	return os.Remove(ConfigFile)
}
