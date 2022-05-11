package red

import "github.com/go-redis/redis/v8"

func GetRedisConnection() *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:     "localhost:8989",
		Password: "",
		DB:       0,
	})

	return connection
}
