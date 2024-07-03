package config

import (
	"encoding/json"
	"os"
)

var ConfigFile *config

type config struct {
	Port int    `json:"Port"` // порт для сервера
	Root string `json:"Root"` // корневая папка
}

func GetConfigData(file string) error {
	if ConfigFile == nil {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		var config config
		err = json.Unmarshal(data, &config)
		if err != nil {
			return err
		}

		ConfigFile = &config
		return nil

	}
	return nil
}
