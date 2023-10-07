package main

import (
	"log/slog"

	"github.com/Coflnet/sky-mayor/internal/metrics"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/Coflnet/sky-mayor/internal/rest"
	"github.com/Coflnet/sky-mayor/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Warn("error loading .env file", "err", err)
	}

	err = mongo.Init()
	if err != nil {
		slog.Error("error connecting to database", "err", err)
		panic(err)
	}
	defer mongo.Disconnect()

	go func() {
		err := rest.Init()
		slog.Error("error starting api", "err", err)
		panic(err)
	}()

	go metrics.Init()
	usecase.StartFetch()
}
