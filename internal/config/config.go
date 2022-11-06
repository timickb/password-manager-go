package config

import "os"

type Config struct {
	Path string `json:"path"`
}

func New() (*Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return &Config{
		Path: homedir,
	}, nil
}
