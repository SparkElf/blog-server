package routes

import (
	"github.com/SparkElf/blog-server/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	server := gin.Default()
	initAuth(server)
	initNormal(server)
	server.Run(config.HttpPort)
}
