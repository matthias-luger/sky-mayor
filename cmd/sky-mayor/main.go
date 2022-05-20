package main

import (
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/Coflnet/sky-mayor/internal/rest"
	"github.com/Coflnet/sky-mayor/internal/usecase"
	"github.com/rs/zerolog/log"
)

func main() {
	err := mongo.Init()

	if err != nil {
		log.Fatal().Err(err).Msg("error while initializing database")
	}

	defer mongo.Disconnect()

	go rest.Init()

	usecase.StartFetch()
}
