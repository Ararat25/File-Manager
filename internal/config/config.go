package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port int `json:"Port"` // порт для сервера
}

func GetConfigData(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
