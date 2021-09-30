package routes

import (
	v1 "github.com/SparkElf/blog-server/api/v1"
	"github.com/gin-gonic/gin"
)

func initNormal(server *gin.Engine) {
	normal := server.Group("api/v1")
	normal.GET("user", v1.GetUser)
}
