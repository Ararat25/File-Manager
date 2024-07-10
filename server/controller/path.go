package controller

import (
	"RBS-Task-3/server/config"
	"RBS-Task-3/server/pkg/fileProperty"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Response структура для записи ответа от сервера
type Response struct {
	Status    int                 `json:"Status"`    // статус ответа (200 - успешно, остальное - ошибка)
	Error     string              `json:"Error"`     // строка с ошибкой
	TimeSpent int                 `json:"TimeSpent"` // затраченное на запрос время
	Root      string              `json:"Root"`      // корневая папка
	Files     []fileProperty.File `json:"Files"`     // слайс со структурами File
}

// RequestToPhp структура для отправки данных на php сервер
type RequestToPhp struct {
	Root      string `json:"root"`      // путь до папки
	Size      int    `json:"size"`      // размер папки
	TimeSpent int    `json:"timeSpent"` // затраченное на запрос время
}

type ResponseFromPhp struct {
	Status  string `json:"status"`  // статус ответа от php сервера
	Message string `json:"message"` // строка с ошибкой
}

// writeResponse записывает в ответ структуру Response с определенными значениями
func writeResponse(w http.ResponseWriter, root string, reqUrl string, status int, error string, files []fileProperty.File, elapsed int) {
	resp, err := json.Marshal(Response{
		Status:    status,
		Error:     error,
		Files:     files,
		TimeSpent: elapsed,
		Root:      root,
	})
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}

	w.WriteHeader(status)

	w.Write(resp)
}

// sendRequest отправляет данные на php сервер
func sendRequest(reqUrl string, url string, output []fileProperty.File, root string, elapsed int) {
	var fullSize int64
	for _, file := range output {
		fullSize += file.ByteSize
	}

	data := RequestToPhp{
		Root:      root,
		Size:      int(fullSize),
		TimeSpent: elapsed,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	respPhp, err := client.Do(req)
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}
	defer respPhp.Body.Close()

	var responseFromPhp ResponseFromPhp
	var buf bytes.Buffer

	_, err = buf.ReadFrom(respPhp.Body)
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}

	err = json.Unmarshal(buf.Bytes(), &responseFromPhp)
	if err != nil {
		log.Printf("%v %v", reqUrl, err.Error())
		return
	}

	if responseFromPhp.Status != "success" {
		log.Printf("%v status: %v %v", reqUrl, responseFromPhp.Status, responseFromPhp.Message)
		return
	}

	log.Printf("%v status: %v", reqUrl, responseFromPhp.Status)
}

// PathHandle обрабатывает HTTP-запросы для получения свойств файлов по указанному пути с возможностью сортировки
func PathHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if config.ConfigFile == nil {
		log.Printf("Ошибка: Не удалось загрузить config")

		writeResponse(w, "", r.URL.String(), http.StatusInternalServerError, "Ошибка: Не удалось загрузить config", nil, 0)
		return
	}
	conf := *config.ConfigFile

	url := fmt.Sprintf("%s:%v%v", conf.UrlPhp, conf.PortPhp, conf.PathPhp)

	root := r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	sort = strings.ToLower(sort)

	if root == "" {
		root = conf.Root
	}

	if !strings.HasPrefix(root, conf.Root) {
		log.Printf("%v %v", r.URL, "Предел глубины")

		writeResponse(w, "", r.URL.String(), http.StatusBadRequest, "Это предел", nil, 0)
		return
	}

	if sort == "" {
		sort = fileProperty.ASC
	}

	if !(sort == fileProperty.ASC || sort == fileProperty.DESC) {
		log.Printf("%v Ошибка: флаг сорт не может быть с таким значением.", r.URL)

		writeResponse(w, "", r.URL.String(), http.StatusBadRequest, "Ошибка: флаг сорт не может быть с таким значением.", nil, 0)
		return
	}

	output, err := fileProperty.OutputFileProperty(root, sort)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		writeResponse(w, "", r.URL.String(), http.StatusInternalServerError, err.Error(), nil, 0)
		return
	}

	elapsed := time.Since(start)

	go sendRequest(r.URL.String(), url, output, root, int(elapsed.Milliseconds()))

	w.Header().Add("Content-Type", "application/json")

	writeResponse(w, root, r.URL.String(), http.StatusOK, "", output, int(elapsed.Milliseconds()))
}
