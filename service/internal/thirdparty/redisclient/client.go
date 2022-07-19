package redisclient

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
)

func NewRedisClient(in digIn) IRedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     in.ServiceConf.GetRedisServiceConfig().Address,
			Password: in.ServiceConf.GetRedisServiceConfig().Password,
			DB:       in.ServiceConf.GetRedisServiceConfig().DB,
		}),
	}
}

type digIn struct {
	dig.In

	ServiceConf config.IServiceConfig
}
