package session

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/emirrcaglar/go-url-shortener/types"
)

type Cfg struct {
	LoggedIn    bool        `json:"loggedIn"`
	CurrentUser *types.User `json:"currentUser"`
}

var ConfigFile string

func init() {
	home, _ := os.UserHomeDir()
	ConfigFile = filepath.Join(home, ".urlshortener", "config.json")
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
	return cfg, err
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
