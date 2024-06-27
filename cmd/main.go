package main

import (
	"RBS-Task-3/controller"
	"RBS-Task-3/internal"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	internal.LoadEnvVariables()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := getServer()

	err := runServer(ctx, server)
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

	err := internal.Recover(ctx, server)
	if err != nil {
		return err
	}

	return nil
}

// getServer возвращает сервер с определенными параметрами
func getServer() *http.Server {
	port := os.Getenv("SERVER_PORT")
	listenAddr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/path", controller.PathHandle)

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	return srv
}
