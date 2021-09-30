package model

import (
	"sync"

	"github.com/SparkElf/blog-server/config"
	"github.com/SparkElf/blog-server/middleware/redis"
)

var U User

func Init() {
	U.Username = config.Username
	U.Github = &Github{Mu: &sync.RWMutex{}}
	U.TotalStars, _ = redis.RedisDB.Get(redis.Ctx, TOTAL_STARS).Int()
	InitDb()
}
