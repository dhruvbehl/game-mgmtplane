package mgmtplane

import (
	"context"
	"net/http"
	"strconv"

	pbengine "github.com/dhruvbehl/game-apis/game-engine/v1"
	pbhighscore "github.com/dhruvbehl/game-apis/game-highscore/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// type for game client
type gameResource struct {
	gameClient   pbhighscore.GameClient
	engineClient pbengine.GameEngineClient
}

func NewGameResource(gameClient pbhighscore.GameClient, engineClient pbengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:   gameClient,
		engineClient: engineClient,
	}
}

func NewGrpcGameServiceClient(serverAddress string) (pbhighscore.GameClient, error) {
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("failed to dial")
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddress)
	}
	if connection == nil {
		log.Error().Msg("connection initiated from mgmtplane is nil for game-highscore")
	}

	client := pbhighscore.NewGameClient(connection)

	return client, nil
}

func NewGrpcEngineServiceClient(serverAddress string) (pbengine.GameEngineClient, error) {
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("failed to dial")
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddress)
	}
	if connection == nil {
		log.Error().Msg("connection initiated from mgmtplane is nil for game-engine")
	}

	client := pbengine.NewGameEngineClient(connection)

	return client, nil
}

func (gr *gameResource) SetHighScore(ctx *gin.Context) {

	highScoreString := ctx.Param("highscore")
	highScore, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert highscore to float")
	}
	_, err = gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{HighScore: highScore})
	if err != nil {
		log.Error().Err(err).Msg("failed to set highscore to service game-highscore")
	}
}

func (gr *gameResource) GetHighScore(ctx *gin.Context) {
	response, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("error while getting highscore")
	}
	ctx.JSONP(http.StatusOK, gin.H{"highscore": response.GetHighScore()})
}

func (gr *gameResource) GetSize(ctx *gin.Context) {
	response, err := gr.engineClient.GetSize(context.Background(), &pbengine.GetSizeRequest{})
	if err != nil {
		log.Error().Err(err).Msg("error while getting size")
	}
	ctx.JSONP(http.StatusOK, gin.H{"size": response.GetSize()})
}

func (gr *gameResource) SetScore(ctx *gin.Context) {
	scoreString := ctx.Param("score")
	score, err := strconv.ParseFloat(scoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert score to float")
	}

	_, err = gr.engineClient.SetScore(context.Background(), &pbengine.SetScoreRequest{Score: score})
	if err != nil {
		log.Error().Err(err).Msg("failed to set size to service game-engine")
	}
}
