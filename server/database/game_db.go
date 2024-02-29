package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/redis/go-redis/v9"
)

func (redisDBCtrl *RedisDBController) GetGameDetails(ctx context.Context, gameId string) (*models.Game, error) {
	val, err := redisDBCtrl.DB.Get(ctx, gameId).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("key not found")
	} else if err != nil {
		return nil, fmt.Errorf("error getting key: %w", err)
	}
	if val == "" {
		return nil, errors.New("key not found")
	}
	var game models.Game
	err = json.Unmarshal([]byte(val), &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (redisDBCtrl *RedisDBController) SaveGameDetails(ctx context.Context, gameId string, modifiedGame *models.Game) error {
	gameBytes, err := json.Marshal(modifiedGame)
	if err != nil {
		return err
	}
	err = redisDBCtrl.DB.Set(ctx, gameId, gameBytes, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
