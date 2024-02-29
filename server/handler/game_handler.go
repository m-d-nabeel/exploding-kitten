package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/m-d-nabeel/exploding-kittens/database"
	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/m-d-nabeel/exploding-kittens/utils"
	"github.com/redis/go-redis/v9"
)

func (apiCfgHandlr *apiConfigHandler) HandlerGetAllGameDetails(w http.ResponseWriter, r *http.Request, userData *models.User) {
	redisDBCtrl := database.RedisDBController{
		ApiConfig: apiCfgHandlr.ApiConfig,
	}
	activeGameId := "active_game:" + userData.ID.String()
	highestScoreGameId := "highest_score_game:" + userData.ID.String()
	activeGame, err := redisDBCtrl.GetGameDetails(r.Context(), activeGameId)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			utils.RespondWithError(w, http.StatusNotFound, "user not found")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	highestScoreGame, err := redisDBCtrl.GetGameDetails(r.Context(), highestScoreGameId)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			utils.RespondWithError(w, http.StatusNotFound, "user not found")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"active_game":        activeGame,
		"highest_score_game": highestScoreGame,
	},
	)
}

func (apiCfgHandlr *apiConfigHandler) HandlerGetTopScores(w http.ResponseWriter, r *http.Request, userData *models.User) {
	redisDBCtrl := database.RedisDBController{
		ApiConfig: apiCfgHandlr.ApiConfig,
	}
	topScores, err := redisDBCtrl.GetTop10Scorers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, topScores)
}

func (apiCfgHandlr *apiConfigHandler) HandlerGameMove(w http.ResponseWriter, r *http.Request, userData *models.User) {
	redisDBCtrl := database.RedisDBController{
		ApiConfig: apiCfgHandlr.ApiConfig,
	}
	activeGameId := "active_game:" + userData.ID.String()
	activeGame, err := redisDBCtrl.GetGameDetails(r.Context(), activeGameId)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			utils.RespondWithError(w, http.StatusNotFound, "user not found")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	cardId := chi.URLParam(r, "cardId")
	if cardId == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "cardId is required")
		return
	}
	var cardType models.CardType
	for _, card := range activeGame.Deck {
		if card.ID.String() == cardId {
			switch card.Type {
			case models.DiffuseCard:
				cardType = models.DiffuseCard
			case models.ExplodingCard:
				cardType = models.ExplodingCard
			case models.ShuffleCard:
				cardType = models.ShuffleCard
			case models.CatCard:
				cardType = models.CatCard
			}
			break
		}
	}

	modifiedGame, err := utils.GetGameResultForMove(activeGame, &cardType)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = redisDBCtrl.SaveGameDetails(r.Context(), activeGameId, modifiedGame)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, modifiedGame)
}
