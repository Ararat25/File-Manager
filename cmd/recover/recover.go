package recover

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const shutdownTimeout = 3 * time.Second

// Recover плавно завершает работу сервера
func Recover(ctx context.Context, srv *http.Server) error {
	<-ctx.Done()

	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := srv.Shutdown(ctxShutdown)
	if err != nil {
		return fmt.Errorf("shutdown: %v", err)
	}

	log.Println("Завершение работы сервера")
	return nil
}
