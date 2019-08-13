package handlers

import (
	"github.com/gin-gonic/gin"
)

// Respond is default api return value
func Respond(c *gin.Context, code int, err string, data ...map[string]interface{}) {
	c.JSON(code, gin.H{
		"code":  code,
		"data":  data,
		"error": err,
	})
}

// Error is default api return when got some error
func Error(c *gin.Context, code int, err string) {
	Respond(c, code, err)
}
