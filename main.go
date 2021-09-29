package main

import (
	"time"

	"github.com/SparkElf/blog-server/config"
	"github.com/SparkElf/blog-server/middleware/redis"
	"github.com/SparkElf/blog-server/model"
	"github.com/SparkElf/blog-server/service/github"
)

func main() {
	launch()
	time.Sleep(5 * time.Minute)
}

func launch() {
	config.Init()
	redis.Init()
	model.Init()
	github.Init()
}
