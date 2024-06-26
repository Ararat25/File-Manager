package main

import (
	"RBS-Task-3/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/path", controller.PathHandle)

	log.Println("Запуск сервера")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
		return
	}
}
