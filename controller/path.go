package controller

import (
	"RBS-Task-3/pkg/fileProperty"
	"encoding/json"
	"log"
	"net/http"
)

const (
	ASK  = "ASK"
	DESC = "DESC"
)

func PathHandle(w http.ResponseWriter, r *http.Request) {
	root := r.URL.Query().Get("root")
	sort := r.URL.Query().Get("sort")

	if root == "" {
		log.Printf("%v Ошибка: пропущены нужные флаги.", r.URL)
		http.Error(w, "Ошибка: пропущены нужные флаги.", http.StatusBadRequest)
		return
	}

	if sort == "" {
		sort = ASK
	}

	if !(sort == ASK || sort == DESC) {
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

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}
