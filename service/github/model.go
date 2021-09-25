package github

import (
	"fmt"
	"forestlightelf/blog-server/middleware/redis"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"sync"
)

const (
	TOTAL_STARS     = "total_stars"
	TOTAL_FOLLOWERS = "total_followers"
)

var (
	TotalStars     int
	TotalFollowers int
)
var mutex sync.RWMutex

func ReadTotalStars() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return TotalStars
}
func WriteTotalStars(num int) {
	mutex.Lock()
	TotalStars = num
	redis.RedisDB.Set(redis.Ctx, TOTAL_STARS, TotalStars, 0) //Zero expiration means the key has no expiration time
	mutex.Unlock()
}
func GetUserTotalStars(username string) (total int) {
	client := resty.New()
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := client.R().
		SetAuthToken("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"). //请求头变为 Authorization Token token
		Get(url)
	if err != nil {
		panic(err)
	}
	value := gjson.Get(resp.String(), "#.stargazers_count").Array()
	total = 0
	for i := range value {
		println("valueArray:", value[i].Int())
		total += int(value[i].Int())
	}
	return total
}
