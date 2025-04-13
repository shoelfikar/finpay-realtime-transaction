package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetCache(redisClient *redis.Client, redisKey string) (string, error) {
    val, err := redisClient.Get(ctx, redisKey).Result()
    if err != nil {
        return "", err
    }
    return val, nil
}

func SetCache(redisClient *redis.Client, redisKey string, data string) error {
    return redisClient.Set(ctx, redisKey, data, 10*time.Minute).Err()
}
