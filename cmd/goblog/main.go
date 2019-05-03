package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var app *gin.Engine
var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbHost := getEnv("DBHOST", "localhost")
	dbUser := getEnv("DBUSER", "")
	dbPort := getEnv("DBPORT", "")
	dbPass := getEnv("DBPASS", "")
	dbName := getEnv("DBNAME", "")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}

func main() {

	var hostIP = getEnv("HOST_IP", "127.0.0.1")
	var port = getEnv("PORT", "8080")
	var debug = getEnv("DEBUG", "true")

	app = gin.Default()

	if debug == "false" {
		gin.SetMode(gin.ReleaseMode)
	}

	initializeRoutes(app)

	app.Run(fmt.Sprintf("%s:%s", hostIP, port))

}

func getEnv(env, fallback string) string {
	e := os.Getenv(env)

	if e == "" {
		return fallback
	}

	return e
}

func initializeRoutes(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

}
