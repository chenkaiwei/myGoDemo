package redis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis.chenkaiwei.com:6379",
	Password: "Ckw_1988", //
	DB:       2,          //
})

var ctx = context.Background()
