package github

import (
	"fmt"
	"forestlightelf/blog-server/middleware/redis"
	"time"
)

func Init() {
	//启动时先从redis中读取数据
	TotalStars, _ = redis.RedisDB.Get(redis.Ctx, TOTAL_STARS).Int()
	println("b")
	go startTask()
}

func startTask() {
	fmt.Println("d")
	//每隔1秒向github请求一次数据，超时时间为5秒
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	fmt.Println("c")
	for {
		select {
		case <-t.C:
			fmt.Println("ready")
			time.Sleep(5 * time.Second)
			fmt.Println("handle")
		}
	}
	//totalstars := GetUserTotalStars(config.Username)

}
