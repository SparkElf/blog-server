package v1

import (
	"net/http"

	"github.com/SparkElf/blog-server/model"
	"github.com/gin-gonic/gin"
)

//查询作者信息
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": model.U,
	})
}
