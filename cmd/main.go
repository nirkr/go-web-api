package main

import (
	"github.com/rs/zerolog/log"
	config "simple_bank"
	"simple_bank/api"
	"simple_bank/internal/controllers"
	"simple_bank/internal/repository"
)

func main() {
	// set config envs
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed loading configuration")
		return
	}

	// init DB
	err = repository.InitPGClient(loadConfig.DBSource)
	if err != nil {
		log.Error().Err(err).Msgf("failed to connect to postgres DB")
	}
	log.Info().Msg("Postgres DB was initiated successfully")

	// init ControllerStore
	store := api.ControllerStore{
		AccountController: controllers.NewAccountController(),
	}
	// init Server
	server := api.NewServer(store)
	err = server.Run(loadConfig.ServerPort)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
}
