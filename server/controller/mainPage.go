package controller

import (
	"html/template"
	"log"
	"net/http"
)

// MainPage обработчик, который возвращает html страницу
func MainPage(res http.ResponseWriter, req *http.Request) {
	templ, err := template.ParseFiles("./client/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(res, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
