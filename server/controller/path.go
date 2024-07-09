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
	Status int                 `json:"Status"`
	Error  string              `json:"Error"`
	Files  []fileProperty.File `json:"Files"`
}

// RequestToPhp структура для отправки данных на php сервер
type RequestToPhp struct {
	Root      string `json:"root"`
	Size      int    `json:"size"`
	TimeSpent int    `json:"timeSpent"`
}

// writeResponse записывает в ответ структуру Response с определенными значениями
func writeResponse(w http.ResponseWriter, status int, error string, files []fileProperty.File) {
	resp, _ := json.Marshal(Response{
		Status: status,
		Error:  error,
		Files:  files,
	})

	w.WriteHeader(status)

	w.Write(resp)
}

// sendRequest отправляет данные на php сервер
func sendRequest(w http.ResponseWriter, r *http.Request, url string, output []fileProperty.File, root string, elapsed time.Duration) {
	var fullSize int64
	for _, file := range output {
		fullSize += file.ByteSize
	}

	data := RequestToPhp{
		Root:      root,
		Size:      int(fullSize),
		TimeSpent: int(elapsed.Milliseconds()),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	respPhp, err := client.Do(req)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer respPhp.Body.Close()
}

// PathHandle обрабатывает HTTP-запросы для получения свойств файлов по указанному пути с возможностью сортировки
func PathHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if config.ConfigFile == nil {
		log.Printf("Ошибка: Не удалось загрузить config")

		writeResponse(w, http.StatusInternalServerError, "Ошибка: Не удалось загрузить config", nil)
		return
	}
	conf := *config.ConfigFile

	url := fmt.Sprintf("%s:%v%v", conf.UrlPhp, conf.PortPhp, conf.PathPhp)

	root := conf.Root + r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	sort = strings.ToLower(sort)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if root == "" {
		log.Printf("%v Ошибка: пропущены нужные флаги.", r.URL)

		writeResponse(w, http.StatusBadRequest, "Ошибка: пропущены нужные флаги.", nil)
		return
	}

	if sort == "" {
		sort = fileProperty.ASC
	}

	if !(sort == fileProperty.ASC || sort == fileProperty.DESC) {
		log.Printf("%v Ошибка: флаг сорт не может быть с таким значением.", r.URL)

		writeResponse(w, http.StatusBadRequest, "Ошибка: флаг сорт не может быть с таким значением.", nil)
		return
	}

	output, err := fileProperty.OutputFileProperty(root, sort)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	elapsed := time.Since(start)

	go sendRequest(w, r, url, output, root, elapsed)

	w.Header().Add("Content-Type", "application/json")

	writeResponse(w, http.StatusOK, "", output)
}
