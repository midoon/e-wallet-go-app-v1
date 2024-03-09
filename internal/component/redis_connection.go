package component

import (
	"context"
	"fmt"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/redis/go-redis/v9"
)

func GetRedisConnection(cnf *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Addr,
		Password: cnf.Redis.Password,
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	return rdb
}
