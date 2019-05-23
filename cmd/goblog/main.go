package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/soerjadi/GoBlog/database"
	h "github.com/soerjadi/GoBlog/handlers"
	"github.com/soerjadi/GoBlog/utils"
)

var app *gin.Engine
var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.InitDB()

}

func main() {

	var hostIP = utils.GetEnv("HOST_IP", "127.0.0.1")
	var port = utils.GetEnv("PORT", "8080")
	var debug = utils.GetEnv("DEBUG", "true")

	app = gin.Default()

	if debug == "false" {
		gin.SetMode(gin.ReleaseMode)
	}

	initializeRoutes(app)

	app.Run(fmt.Sprintf("%s:%s", hostIP, port))

}

func initializeRoutes(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	app.GET("/user/info", h.UserInfo)

}
