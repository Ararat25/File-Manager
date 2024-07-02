package main

import (
	"RBS-Task-3/cmd/recover"
	"RBS-Task-3/controller"
	"RBS-Task-3/internal/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server, err := getServer()
	if err != nil {
		log.Fatal(err)
	}

	err = runServer(ctx, server)
	if err != nil {
		log.Fatal(err)
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
func getServer() (*http.Server, error) {
	conf, err := config.GetConfigData("config.json")
	if err != nil {
		return nil, err
	}

	listenAddr := fmt.Sprintf(":%v", conf.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/path", controller.PathHandle)
	mux.HandleFunc("/fs", controller.MainPage)

	fileServer := http.FileServer(http.Dir("./view/static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	return srv, nil
}
