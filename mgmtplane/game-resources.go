package mgmtplane

import (

	pbengine "github.com/dhruvbehl/game-apis/game-engine/v1"
	pbhighscore "github.com/dhruvbehl/game-apis/game-highscore/v1"
	"google.golang.org/grpc"
	"github.com/rs/zerolog/log"
)

// type for game client
type gameResource struct {
	gameClient pbhighscore.GameClient
	engineClient pbengine.GameEngineClient
}

func NewGameResource(gameClient pbhighscore.GameClient, engineClient pbengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient: gameClient,
		engineClient: engineClient,
	}
}

func NewGrpcGameServiceClient(serverAddress string) (pbhighscore.GameClient, error) {
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial")
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddress)
	}
	if connection == nil {
		log.Fatal().Msg("connection initiated from mgmtplane is nil for game-highscore")
	}

	client := pbhighscore.NewGameClient(connection)

	return client, nil
}

func NewGrpcEngineServiceClient(serverAddress string) (pbengine.GameEngineClient, error) {
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial")
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddress)
	}
	if connection == nil {
		log.Fatal().Msg("connection initiated from mgmtplane is nil for game-engine")
	}

	client := pbengine.NewGameEngineClient(connection)

	return client, nil
}