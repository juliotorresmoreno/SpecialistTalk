package services

import (
	"github.com/go-redis/redis"
	"github.com/juliotorresmoreno/SpecialistTalk/configs"
)

var poolRedis *redis.Client

func GetPoolRedis() *redis.Client {
	if poolRedis == nil {
		conf := configs.GetConfig()
		poolRedis = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
			PoolSize: conf.Redis.PoolSize,
		})
	}

	return poolRedis
}

func NewRedis() *redis.Client {
	conf := configs.GetConfig()
	redisCli := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
		PoolSize: 1,
	})
	return redisCli
}
