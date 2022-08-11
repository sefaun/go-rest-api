package connections

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

type Redis struct {
	RedisConnection *redis.Client
}

func (con *Redis) CreateRedisConnection() {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	options := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	con.RedisConnection = redis.NewClient(options)
}

func (con *Redis) GetSetTesting(data string) (string, error) {
	_, err := con.RedisConnection.Set("sefa", data, 0).Result()

	if err != nil {
		fmt.Println("Redis Set Error", err)
	}

	return con.RedisConnection.Get("sefa").Result()
}
