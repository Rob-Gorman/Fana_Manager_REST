package main

import (
	"context"
	"fmt"
	"manager/cmd/api"
	"manager/dev"
	"manager/utils"
	"os"
	"os/signal"
	"time"
)

func main() {
	utils.InitLoggers(nil, nil)
	utils.LoadDotEnv()
	app := api.NewApp()
	dev.RefreshSchema(app.H.DM)
	fmt.Println("Connected to postgres!")

	srv := app.NewServer()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			utils.ErrLog.Falalf("%v", err)
		}
	}()
	// utils.Shutdown(context.Background(), srv) // os.Kill not valid here
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan // blocks
	utils.InfoLog.Printf("Shutting down server %v", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
