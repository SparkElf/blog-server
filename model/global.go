package model

import (
	"github.com/SparkElf/blog-server/config"
	"github.com/SparkElf/blog-server/middleware/redis"
)

var U User

func Init() {
	U.Username = config.Username
	U.TotalStars, _ = redis.RedisDB.Get(redis.Ctx, TOTAL_STARS).Int()
}
