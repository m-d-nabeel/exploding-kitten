package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/redis/go-redis/v9"
)

func (redisDBCtrl *RedisDBController) CreateUser(ctx context.Context, userData models.User) error {
	data, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	key := "user:" + userData.Username
	err = redisDBCtrl.DB.Get(ctx, key).Err()
	if err == nil {
		return errors.New("key already exists")
	}
	txn := redisDBCtrl.DB.TxPipeline()
	err = txn.SetNX(ctx, key, string(data), 0).Err()
	if err != nil {
		txn.Discard()
		return err
	}

	activeGameId := "active_game:" + userData.ID.String()
	highestScoreGameId := "highest_score_game:" + userData.ID.String()
	activeGame := models.NewGame(userData.ID, models.ActiveGame)
	highestScoreGame := models.NewGame(userData.ID, models.FinishedGame)
	mrslActGame, err := json.Marshal(activeGame)
	if err != nil {
		return err
	}
	mrslHghstScrGame, err := json.Marshal(highestScoreGame)
	if err != nil {
		return err
	}
	err = txn.SetNX(ctx, activeGameId, mrslActGame, 0).Err()
	if err != nil {
		txn.Discard()
		return err
	}
	err = txn.SetNX(ctx, highestScoreGameId, mrslHghstScrGame, 0).Err()
	if err != nil {
		txn.Discard()
		return err
	}

	if _, err := txn.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (redisDBCtrl *RedisDBController) GetUserDetail(ctx context.Context, key string) (*models.User, error) {
	value, err := redisDBCtrl.DB.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("key not found")
	} else if err != nil {
		return nil, fmt.Errorf("error getting key: %w", err)
	}

	if !strings.Contains(key, ":") {
		return nil, errors.New("invalid key")
	}

	modelType := strings.Split(key, ":")[0]
	if modelType != "user" {
		return nil, errors.New("invalid key")
	}
	var data = models.User{}

	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling card: %w", err)
	}

	return &data, nil
}
