package main

import (
	"context"
	"embed"
	"fmt"
	"manager/cmd/api"
	"manager/dev"
	"manager/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed static/*
var static embed.FS

func main() {
	utils.InitLoggers(nil, nil)
	utils.LoadDotEnv()
	app := api.NewApp(static)
	dev.RefreshSchema(app.H.DM)
	fmt.Println("Connected to postgres!")

	srv := app.NewServer()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			utils.ErrLog.Falalf("%v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan // blocks
	utils.InfoLog.Printf("Shutting down server %v", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
