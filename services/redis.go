package services

import (
	"github.com/go-redis/redis"
	"github.com/juliotorresmoreno/freelive/configs"
)

func NewRedis() *redis.Client {
	conf := configs.GetConfig()
	redisCli := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	return redisCli
}
