package main

import (
	"github.com/SparkElf/blog-server/config"
	"github.com/SparkElf/blog-server/middleware/redis"
	"github.com/SparkElf/blog-server/model"

	"github.com/SparkElf/blog-server/routes"
	"github.com/SparkElf/blog-server/service"
)

func main() {
	launch()

}

func launch() {
	config.Init()
	redis.Init()
	model.Init()
	service.Init()
	routes.InitRouter()
}
