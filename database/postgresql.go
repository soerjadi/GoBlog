package database

import (
	"database/sql"
	"fmt"

	// import util
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soerjadi/GoBlog/utils"
)

func InitDB() *sql.DB {

	if err := godotenv.Load("../.env"); err == nil {
		panic("Error loading .env file")
	}

	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbUser := utils.GetEnv("DB_USER", "")
	dbPort := utils.GetEnv("DB_PORT", "")
	dbPass := utils.GetEnv("DB_PASS", "")
	dbName := utils.GetEnv("DB_NAME", "")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}

func InitTestDB() *sql.DB {

	if err := godotenv.Load("../.env"); err == nil {
		panic("Error loading .env file")
	}

	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbUser := utils.GetEnv("DB_USER", "")
	dbPort := utils.GetEnv("DB_PORT", "")
	dbPass := utils.GetEnv("DB_PASS", "")
	dbName := utils.GetEnv("DB_TEST", "")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Println(connStr)

	dbTest, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = dbTest.Ping()
	if err != nil {
		panic(err)
	}

	return dbTest
}
