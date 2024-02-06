package server

import (
	"context"
	"log"
	"time"

	"iharacee/config"

	"github.com/redis/go-redis/v9"
)

type RedisInstance struct {
	Client *redis.Client
	ctx    context.Context
}

var Redis RedisInstance

// Source https://go.dev/tour/methods/1
func (db RedisInstance) Get(key string) string {
	val, err := db.Client.Get(db.ctx, key).Result()
	if err == redis.Nil {
		val = ""
	} else if err != nil {
		val = ""
	}
	return val
}
func (db RedisInstance) Del(keys ...string) error {
	err := db.Client.Del(db.ctx, keys...).Err()
	return err
}
func (db RedisInstance) Set(key string, value interface{}) error {
	err := db.Client.Set(db.ctx, key, value, 0).Err()
	return err
}
func (db RedisInstance) Expire(key string, expiration time.Duration) error {
	err := db.Client.Expire(db.ctx, key, expiration).Err()
	return err
}

// Source: https://redis.io/docs/clients/go/
func init() {
	opt, err := redis.ParseURL(config.REDIS_URI)

	if err != nil {
		log.Println("Error parsing redis url: " + err.Error())
		log.Fatalln(err)
	} else {
		log.Println("Redis successfully connected...")
	}

	ctx := context.Background()
	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}

	Redis = RedisInstance{
		Client: client,
		ctx:    ctx,
	}
}
