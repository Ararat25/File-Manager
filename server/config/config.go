package config

import (
	"encoding/json"
	"os"
)

// ConfigFile глобальная переменная для получения данных из config.json файла
var ConfigFile *config

// config cтруктура config.json файла
type config struct {
	Port    int    `json:"Port"`     // порт для сервера
	Root    string `json:"Root"`     // корневая папка
	UrlPhp  string `json:"Url_php"`  // url до php
	PortPhp int    `json:"Port_php"` // порт для php сервера
	PathPhp string `json:"Path_php"` // путь до php обработчика
}

// UploadConfigData считывает config.json файл и записывает значения в структуру ConfigFile
func UploadConfigData(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var conf config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}

	ConfigFile = &conf
	return nil
}
