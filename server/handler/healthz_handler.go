package handler

import (
	"net/http"

	"github.com/m-d-nabeel/exploding-kittens/database"
	"github.com/m-d-nabeel/exploding-kittens/utils"
)

type HealthzHandler struct {
	Status string `json:"status"`
	Health string `json:"health"`
}

func (apiCfgHandlr *apiConfigHandler) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	resp := HealthzHandler{
		Status: "ok",
		Health: "good",
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (apiCfgHandlr *apiConfigHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	var data []interface{}
	redisDBCtrl := database.RedisDBController{
		ApiConfig: apiCfgHandlr.ApiConfig,
	}
	data, err := redisDBCtrl.GetAll(r.Context())

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}
