package handler

import (
	"github.com/m-d-nabeel/exploding-kittens/config"
	"github.com/redis/go-redis/v9"
)

type apiConfigHandler struct {
	*config.ApiConfig
}

func NewApiConfigHandler(cfg *config.ApiConfig) *apiConfigHandler {
	return &apiConfigHandler{cfg}
}

func (ach *apiConfigHandler) GetDB() *redis.Client {
	return ach.DB
}
