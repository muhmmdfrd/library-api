package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "",          
        DB:       0,            
    })
}

func StoreSession(ctx context.Context, userID uint, token string) error {
    return RedisClient.Set(ctx, string(userID), token, 24*time.Hour).Err()
}

func GetSession(ctx context.Context, userID uint) (string, error) {
    return RedisClient.Get(ctx, string(userID)).Result()
}

func DeleteSession(ctx context.Context, userID uint) error {
    return RedisClient.Del(ctx, string(userID)).Err()
}