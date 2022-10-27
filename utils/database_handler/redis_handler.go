package database_handler

import (
	"github.com/go-redis/redis/v9"
)

func NewRedisDB(option *redis.Options) *redis.Client {

	rdb := redis.NewClient(option)
	return rdb

}
