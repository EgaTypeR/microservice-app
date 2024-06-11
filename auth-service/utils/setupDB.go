package utils

import (
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var RedisClient *redis.Client
var MySqlDB *gorm.DB

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       1,
	})
	RedisClient = client
}

func SetupMysql() {
	dbUrl := os.Getenv("DB_URL")
	database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate()
	MySqlDB = database
}
