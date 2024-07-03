package controller

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(res http.ResponseWriter, req *http.Request) {
	ts, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(res, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
