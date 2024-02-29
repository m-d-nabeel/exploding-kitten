package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/m-d-nabeel/exploding-kittens/database"
	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/m-d-nabeel/exploding-kittens/utils"
)

func (apiCfgHandlr *apiConfigHandler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user := models.NewUser(params.Name, params.Username)
	log.Println("User: ", user)
	redisDBCtrl := database.RedisDBController{
		ApiConfig: apiCfgHandlr.ApiConfig,
	}
	err = redisDBCtrl.CreateUser(r.Context(), user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (apiCfgHandlr *apiConfigHandler) GetUserDetails(w http.ResponseWriter, r *http.Request, user *models.User) {
	utils.RespondWithJSON(w, http.StatusOK, user)
}
