package config

import (
	"encoding/json"
	"os"
)

// ConfigFile глобальная переменная для получения данных из config.json файла
var ConfigFile *config

// config cтруктура config.json файла
type config struct {
	Port int    `json:"Port"` // порт для сервера
	Root string `json:"Root"` // корневая папка
}

// GetConfigData считывает config.json файл и записывает значения в структуру ConfigFile
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
