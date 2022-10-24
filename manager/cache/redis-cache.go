package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"manager/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

// collection of fields
type redisCache struct {
	host    string
	db      int           //index between 0 and 15
	expires time.Duration // expiration time for all elements in cache in seconds
}

// init creates a new instance of redisCache
func NewRedisCache(host string, db int, exp time.Duration) FlagCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

// create a new redis client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.GetEnvVar("REDIS_HOST"), utils.GetEnvVar("REDIS_PORT")),
		Password: utils.GetEnvVar("REDIS_PW"),
		DB:       cache.db,
	})
}

// Implementation of Set - associate flag json to the key
func (cache *redisCache) Set(key string, value interface{}) {
	if cache == nil {
		utils.ErrLog.Printf("redis client not initialized; cannot set cache")
		return
	}

	client := cache.getClient()

	// serialize the flag
	json, err := json.Marshal(value)
	if err != nil {
		utils.ErrLog.Printf("set from redis marshalling error: %v", err)
		return
	}

	// set the key to marshalled data
	err = client.Set(context.Background(), key, json, cache.expires*time.Second).Err()
	if err != nil {
		utils.ErrLog.Printf("cannot write to redis cache: %v", err)
		return
	}
}

// asynchronously flush all keys from cache
func (cache *redisCache) FlushAllAsync() {
	client := cache.getClient()
	defer client.Close()

	if client == nil {
		utils.ErrLog.Printf("redis cache did not initialize; cannot flush cache")
		return
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println(pong)
		utils.ErrLog.Printf("error with cache.getClient() when flushing cache: %v", err)
		return
	}

	client.FlushAllAsync(context.Background())
}
