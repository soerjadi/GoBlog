package user

import (
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "hello world user",
	})
}