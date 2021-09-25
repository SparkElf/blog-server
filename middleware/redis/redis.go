package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Ctx = context.Background()

func Init() { //关键服务用panic
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //redis有分库功能 一般都是用0号数据库
	})
	err := RedisDB.Ping(Ctx).Err()
	if err != nil {
		panic(err)
	}
}
