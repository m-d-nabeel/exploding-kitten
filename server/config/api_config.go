package config

import (
	"github.com/redis/go-redis/v9"
)

type ApiConfig struct {
	DB *redis.Client
}
