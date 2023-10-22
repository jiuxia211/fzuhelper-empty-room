package cache

import (
	"context"
	"fzuhelper-empty-room/pkg/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.Password,
		DB:       constants.RedisDBEmptyRoom,
	})
	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}
