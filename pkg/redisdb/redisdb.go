package redisdb

import (
	"fmt"
	"ms-auth/config"

	"github.com/go-redis/redis"
)

func InitRedis(config *config.Config) (redisClient *redis.Client, err error) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Hostname, config.Redis.Port),
		Password: config.Redis.Password,
	})

	_, err = redisClient.Ping().Result()

	if err != nil {
		return nil, err
	}

	return redisClient, nil

}
