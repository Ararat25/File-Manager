package controller

import (
	"RBS-Task-3/server/config"
	"RBS-Task-3/server/pkg/fileProperty"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func PathHandle(w http.ResponseWriter, r *http.Request) {
	if config.ConfigFile == nil {
		log.Printf("Ошибка: Не удалось загрузить config")
		http.Error(w, "Ошибка: Не удалось загрузить config", http.StatusInternalServerError)
		return
	}
	conf := *config.ConfigFile

	root := conf.Root + r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	sort = strings.ToLower(sort)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if root == "" {
		log.Printf("%v Ошибка: пропущены нужные флаги.", r.URL)
		http.Error(w, "Ошибка: пропущены нужные флаги.", http.StatusBadRequest)
		return
	}

	if sort == "" {
		sort = fileProperty.ASC
	}

	if !(sort == fileProperty.ASC || sort == fileProperty.DESC) {
		log.Printf("%v Ошибка: флаг сорт не может быть с таким значением.", r.URL)
		http.Error(w, "Ошибка: флаг сорт не может быть с таким значением.", http.StatusBadRequest)
		return
	}

	output, err := fileProperty.OutputFileProperty(root, sort)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(output)
	if err != nil {
		log.Printf("%v %v", r.URL, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}
