package publisher

import (
	"context"
	"fmt"
	"manager/configs"
	"manager/utils"

	"github.com/go-redis/redis/v8"
)

type Pub struct {
	Redis *redis.Client
}

func defaultRedisClient() (client *redis.Client, err error) {
	configs.LoadDotEnv()

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
		Password: configs.GetEnvVar("REDIS_PW"),
		DB:       0, // default
	})

	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		return nil, utils.RedisConnErr(err)
	}

	fmt.Printf("\nRedis publisher client connected at %s\n", client.Options().Addr)
	return client, nil
}

func NewDefaultPublisher() (*Pub, error) {
	client, err := defaultRedisClient()
	return &Pub{Redis: client}, err
}

func (p *Pub) PublishContent(data *[]byte, channel string) error {
	if p.Redis == nil {
		return utils.RedisConnErr(nil)
	}

	err := p.Redis.Publish(context.Background(), channel, *data).Err()
	if err != nil {
		return utils.RedisPublishErr(err)
	}

	return nil
}
