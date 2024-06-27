package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Recover плавно завершает работу сервера
func Recover(ctx context.Context, srv *http.Server) error {
	<-ctx.Done()

	err := srv.Shutdown(context.Background())
	if err != nil {
		return fmt.Errorf("shutdown: %v", err)
	}

	log.Println("Завершение работы сервера")
	return nil
}
