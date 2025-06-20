package config

import (
	"log"

	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/global"
	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})
	if _, err := RedisClient.Ping().Result(); err != nil {
		log.Fatalf("redis connect is error: %v", err)
	}
	global.RedisDB = RedisClient
}
