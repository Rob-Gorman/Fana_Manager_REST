package utils

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Shutdown(ctx context.Context, srv *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan // blocks
	ErrLog.Printf("Shutting down server %v", sig)

	ctx, _ = context.WithTimeout(ctx, 30*time.Second)
	srv.Shutdown(ctx)
}
