package model

import (
	"sync"

	"github.com/SparkElf/blog-server/middleware/redis"
)

type Github struct {
	TotalStars     int
	TotalFollowers int
	Mu             sync.RWMutex
}

const (
	TOTAL_STARS     = "total_stars"
	TOTAL_FOLLOWERS = "total_followers"
)

func (this *Github) GetTotalStars() int {
	this.Mu.RLock()
	defer this.Mu.RUnlock()
	return this.TotalStars
}
func (this *Github) SetTotalStars(num int) {
	this.Mu.Lock()
	this.TotalStars = num
	redis.RedisDB.Set(redis.Ctx, TOTAL_STARS, this.TotalStars, 0) //Zero expiration means the key has no expiration time
	this.Mu.Unlock()
}
