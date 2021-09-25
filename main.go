package main

import (
	"forestlightelf/blog-server/config"
	"forestlightelf/blog-server/middleware/redis"
	"forestlightelf/blog-server/service/github"

	"time"
)

func main() {

	t := time.NewTimer(1 * time.Second)
	defer t.Stop()
	end := true
	start := false
	count := 0
	for s := range t.C {
		println(s.Format("2006-01-02 15:04:05"))
		if !end {
			println("超时")
			return
		}
		println("pass")

		go func() {

			if !start {
				t.Reset(5 * time.Second)
				start = true
				end = false
				count += 1
				println(count)
				if count == 5 {
					time.Sleep(10 * time.Second)
				}
				end = true
				start = false
				t.Reset(1 * time.Second)
			}

		}()
	}
}

func launch() {
	config.Init()
	redis.Init()
	github.Init()
	time.Sleep(30 * time.Minute)
}
