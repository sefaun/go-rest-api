package connections

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var RedisConnection *redis.Client

func CreateRedisConnection() {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	options := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	RedisConnection = redis.NewClient(options)

	println("Redis Connected !")
}
