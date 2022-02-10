package main

import (
	"flag"

	"github.com/dhruvbehl/game-mgmtplane/mgmtplane"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	gameHighScoreAddress := flag.String("address-highscore", "localhost:50001", "Address for the game-highscore gRPC service")
	gameEngineAddress := flag.String("address-engine", "localhost:60001", "Address for the game-engine gRPC service")
	httpAddress := flag.String("address", ":8080", "Address for the game-highscore gRPC service")

	flag.Parse()

	gameClient, err := mgmtplane.NewGrpcGameServiceClient(*gameHighScoreAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create client for game-highscore service")
	}
	engineClient, err := mgmtplane.NewGrpcEngineServiceClient(*gameEngineAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create client for  game-engine service")
	}

	gr := mgmtplane.NewGameResource(gameClient, engineClient)

	router := gin.Default()

	router.GET("/geths", gr.GetHighScore)
	router.GET("/seths/:highscore", gr.SetHighScore)
	router.GET("/getsize", gr.GetSize)
	router.GET("/setscore/:score", gr.SetScore)

	if err := router.Run(*httpAddress); err != nil {
		log.Fatal().Err(err).Msg("couldn't start mgmtplane service")
	}

	log.Info().Msgf("Started http server at [%s]", *httpAddress)
}