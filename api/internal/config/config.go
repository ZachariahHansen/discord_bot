package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	PostgresUser     string `json:"POSTGRES_USER"`
	PostgresPassword string `json:"POSTGRES_PASSWORD"`
	PostgresDB       string `json:"POSTGRES_DB"`
	BaseURL          string `json:"base_url"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
