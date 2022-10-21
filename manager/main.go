package main

import (
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/dev"
	"manager/publisher"
	"net/http"
	"os"
)

func main() {
	configs.LoadDotEnv()
	app := api.NewApp()
	dev.RefreshSchema(app.H.DB)
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")

	publisher.CreateRedisClient()

	fmt.Printf("\nServing flag configuration on PORT %s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
