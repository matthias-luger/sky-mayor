package main

import (
	"github.com/Coflnet/sky-mayor/internal/metrics"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/Coflnet/sky-mayor/internal/rest"
	"github.com/Coflnet/sky-mayor/internal/usecase"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err)
	}

	err = mongo.Init()

	if err != nil {
		log.Fatal().Err(err).Msg("error while initializing database")
	}

	defer mongo.Disconnect()

	go func() {
		err := rest.Init()
		log.Fatal().Err(err).Msg("Error initialicing rest service")
	}()

	go metrics.Init()

	usecase.StartFetch()
}
