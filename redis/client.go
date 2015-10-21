package main

import (
	"fmt"

	"gopkg.in/redis.v3"

	"github.com/nothingelsematters7/golang_proto/config"
	"github.com/nothingelsematters7/golang_proto/utils"
)

func main() {
	conf := config.Conf

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.REDIS_HOST, conf.REDIS_PORT),
		Password: conf.REDIS_PASS,
		DB:       conf.REDIS_DB,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", 123, 0).Err()
	utils.FailOnError(err, "Cannot set key to value in Redis")

	val, err := client.Get("key").Result()
	utils.FailOnError(err, "Cannot read key from Redis")
	fmt.Println("key", val)
}
