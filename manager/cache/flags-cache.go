package cache

import (
	"fmt"
	"manager/configs"
	"manager/utils"
	"strconv"
	"time"
)

// FlagCache is implemented by redisCache struct
type FlagCache interface {
	Set(key string, value interface{}) // set an array of
	FlushAllAsync()
}

func InitFlagCache() FlagCache {
	address := fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT"))
	db, err := strconv.Atoi(configs.GetEnvVar("REDIS_DB"))
	if err != nil {
		utils.HandleErr(err, "error parsing REDIS_DB environment value")
		return nil
	}
	expires, err := time.ParseDuration(configs.GetEnvVar("SECS_TO_EXPIRE"))

	if err != nil {
		utils.HandleErr(err, "error parsing SECS_TO_EXPIRE environment value")
		return nil
	}
	return NewRedisCache(address, db, expires)
}
