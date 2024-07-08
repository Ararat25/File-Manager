package controller

import (
	"RBS-Task-3/server/config"
	"RBS-Task-3/server/pkg/fileProperty"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Response структура для записи ответа от сервера
type Response struct {
	Status int                 `json:"Status"`
	Error  string              `json:"Error"`
	Files  []fileProperty.File `json:"Files"`
}

// PathHandle обрабатывает HTTP-запросы для получения свойств файлов по указанному пути с возможностью сортировки
func PathHandle(w http.ResponseWriter, r *http.Request) {
	if config.ConfigFile == nil {
		log.Printf("Ошибка: Не удалось загрузить config")

		resp, _ := json.Marshal(Response{
			Status: 500,
			Error:  "Ошибка: Не удалось загрузить config",
			Files:  nil,
		})

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(resp)
		return
	}
	conf := *config.ConfigFile

	root := conf.Root + r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	sort = strings.ToLower(sort)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if root == "" {
		log.Printf("%v Ошибка: пропущены нужные флаги.", r.URL)

		resp, _ := json.Marshal(Response{
			Status: 400,
			Error:  "Ошибка: пропущены нужные флаги.",
			Files:  nil,
		})

		w.WriteHeader(http.StatusBadRequest)

		w.Write(resp)
		return
	}

	if sort == "" {
		sort = fileProperty.ASC
	}

	if !(sort == fileProperty.ASC || sort == fileProperty.DESC) {
		log.Printf("%v Ошибка: флаг сорт не может быть с таким значением.", r.URL)

		resp, _ := json.Marshal(Response{
			Status: 400,
			Error:  "Ошибка: флаг сорт не может быть с таким значением.",
			Files:  nil,
		})

		w.WriteHeader(http.StatusBadRequest)

		w.Write(resp)
		return
	}

	output, err := fileProperty.OutputFileProperty(root, sort)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		resp, _ := json.Marshal(Response{
			Status: 500,
			Error:  err.Error(),
			Files:  nil,
		})

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(resp)
		return
	}

	resp, err := json.Marshal(Response{
		Status: 200,
		Error:  "",
		Files:  output,
	})

	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())

		resp, _ := json.Marshal(Response{
			Status: 500,
			Error:  err.Error(),
			Files:  nil,
		})

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(resp)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}
