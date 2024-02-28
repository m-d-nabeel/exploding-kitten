package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/m-d-nabeel/exploding-kittens/config"
	"github.com/m-d-nabeel/exploding-kittens/models"
	"github.com/m-d-nabeel/exploding-kittens/utils"
	"github.com/redis/go-redis/v9"
)

type RedisDBController struct {
	*config.ApiConfig
}

func (redisDBCtrl *RedisDBController) Insert(ctx context.Context, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	key := utils.GenerateKeyForRDB(payload)
	err = redisDBCtrl.DB.Get(ctx, key).Err()
	if err == nil {
		return errors.New("key already exists")
	}
	err = redisDBCtrl.DB.SetNX(ctx, key, string(data), 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (redisDBCtrl *RedisDBController) Get(ctx context.Context, key string) (interface{}, error) {
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
	var data interface{}
	switch modelType {
	case "card":
		data = models.Card{}
	case "user":
		data = models.User{}
	case "game":
		data = models.Game{}
	default:
		return nil, errors.New("invalid key")
	}

	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling card: %w", err)
	}

	return data, nil
}

func (redisDBCtrl *RedisDBController) Delete(ctx context.Context, key string) error {
	err := redisDBCtrl.DB.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		return errors.New("key not found")
	} else if err != nil {
		return fmt.Errorf("error getting key: %w", err)
	}
	return nil
}

func (redisDBCtrl *RedisDBController) Update(ctx context.Context, key string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = redisDBCtrl.DB.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		log.Println(err)
		return errors.New("key not found")
	} else if err != nil {
		log.Println(err)
		return fmt.Errorf("error getting key: %w", err)
	}

	return nil
}

func (redisDBCtrl *RedisDBController) GetTop10Scorers(ctx context.Context) ([]models.User, error) {
	var games []models.Game
	keys, err := redisDBCtrl.DB.Keys(ctx, "game:*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		game, err := redisDBCtrl.Get(ctx, key)
		if err != nil {
			return nil, err
		}
		games = append(games, game.(models.Game))
	}

	sort.Slice(games, func(i, j int) bool {
		return games[i].Score > games[j].Score
	})

	var top10Users []models.User
	idx := 0
	for i := 0; i < 10; i++ {
		if i >= len(games) {
			break
		}
		user, err := redisDBCtrl.Get(ctx, "user:"+games[i].UserId.String())

		if err != nil {
			return nil, err
		}
		top10Users[idx] = models.User{
			ID:        user.(models.User).ID,
			Username:  user.(models.User).Username,
			CreatedAt: user.(models.User).CreatedAt,
			UpdatedAt: user.(models.User).UpdatedAt,
			Name:      user.(models.User).Name,
			Score:     user.(models.User).Score,
		}
		idx += 1
	}

	return top10Users, nil
}

func (redisDBCtrl *RedisDBController) GetAll(ctx context.Context) ([]interface{}, error) {
	keys, err := redisDBCtrl.DB.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	var data []interface{}
	for _, key := range keys {
		value, err := redisDBCtrl.DB.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		var d interface{}
		err = json.Unmarshal([]byte(value), &d)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}
