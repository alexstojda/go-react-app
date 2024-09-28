package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-react-app/internal/app"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	config, err := app.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("could not load config")
	}

	server := app.NewServer(config.SPAPath, config.ClientOrigins)

	err = server.StartServer()
	if err != nil {
		log.Fatal().Err(err).Msg("could not start server")
	}
}
