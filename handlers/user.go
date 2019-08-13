package handlers

import (
	"github.com/gin-gonic/gin"
)

// UserInfo to get info from specific user
func UserInfo(c *gin.Context) {
	//var rv = Response{message: "hello world", errorStr: "aoe"}

	// JSONRV, _ := json.Marshal(rv)
	Respond(c, 200, "", gin.H{
		"message": "hello",
	})
}

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/info", UserInfo)
	}
}
