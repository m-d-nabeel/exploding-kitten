package handler

import (
	"errors"
	"net/http"

	"github.com/m-d-nabeel/exploding-kittens/database"
	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/m-d-nabeel/exploding-kittens/utils"
	"github.com/redis/go-redis/v9"
)

func (apiCfgHandlr *apiConfigHandler) GetAllGameDetails(w http.ResponseWriter, r *http.Request, userData *models.User) {
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
