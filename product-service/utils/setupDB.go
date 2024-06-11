package utils

import (
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitiateDB() {
	dbUrl := os.Getenv("DB_URL")
	database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate()
	DB = database
}

var RedisClient *redis.Client

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	RedisClient = client
}
