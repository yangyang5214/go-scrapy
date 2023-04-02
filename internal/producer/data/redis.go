package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-scrapy/conf"
)

func NewRedis(conf *conf.ServerConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})
	if redisClient.Ping(context.Background()).Err() != nil {
		panic("redis connect fail")
	}
	return redisClient
}
