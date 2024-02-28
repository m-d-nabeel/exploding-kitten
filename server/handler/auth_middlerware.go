package handler

import (
	"log"
	"net/http"

	"github.com/m-d-nabeel/exploding-kittens/auth"
	"github.com/m-d-nabeel/exploding-kittens/database"
	"github.com/m-d-nabeel/exploding-kittens/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, interface{})

func (apiCfgHandlr *apiConfigHandler) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		redisDBCtrl := database.RedisDBController{
			ApiConfig: apiCfgHandlr.ApiConfig,
		}
		data, err := redisDBCtrl.Get(r.Context(), "user:"+apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get user")
			log.Println("Failed to get user: ", err)
			return
		}
		log.Println("Data: ", data)
		handler(w, r, data)
	}
}
