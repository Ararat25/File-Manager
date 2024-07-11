package main

import (
	"RBS-Task-3/server/config"
	"RBS-Task-3/server/controller"
	"RBS-Task-3/server/recover"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func init() {
	err := config.UploadConfigData("./server.config.json")
	if err != nil {
		log.Fatalf("Не удалось загрузить config.json: %v", err)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	defer cancel()

	server := getServer()

	err := runServer(ctx, server)
	if err != nil {
		log.Fatalln(err)
	}
}

// runServer запускает сервер
func runServer(ctx context.Context, server *http.Server) error {
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Запуск сервера: %v", err)
		}
	}()

	log.Printf("Запуск сервера с адресом: %s", server.Addr)

	err := recover.Recover(ctx, server)
	if err != nil {
		return err
	}

	return nil
}

// getServer возвращает сервер с определенными параметрами
func getServer() *http.Server {
	conf := *config.ConfigFile

	listenAddr := fmt.Sprintf(":%v", conf.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/path", controller.PathHandle)
	mux.HandleFunc("/", controller.MainPage)

	fileServer := http.FileServer(http.Dir("./client/static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	return srv
}
