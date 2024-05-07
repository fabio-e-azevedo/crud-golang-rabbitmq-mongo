package rediscache

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var singletonRedisConnection *redis.Client
var redisError error
var once sync.Once

func client() (*redis.Client, error) {
	once.Do(func() {
		url := "redis://redis:6379/0?protocol=3"
		opts, err := redis.ParseURL(url)
		if err != nil {
			redisError = err
		}
		singletonRedisConnection = redis.NewClient(opts)
	})
	return singletonRedisConnection, redisError
}

func Set(name string, data []byte) {
	rdb, err := client()
	if err != nil {
		log.Fatal(err)
	}

	ttl := 1 * time.Minute

	err = rdb.Set(ctx, name, data, ttl).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func Get(name string) (string, error) {
	rdb, err := client()
	if err != nil {
		log.Fatal(err)
	}

	var errResult error

	result, err := rdb.Get(ctx, name).Result()
	if err == redis.Nil {
		errResult = fmt.Errorf("key \"%s\" does not exist", name)
	} else if err != nil {
		errResult = err
	}

	return result, errResult
}
