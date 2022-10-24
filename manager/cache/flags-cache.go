package cache

import (
	"fmt"
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
	address := fmt.Sprintf("%s:%s", utils.GetEnvVar("REDIS_HOST"), utils.GetEnvVar("REDIS_PORT"))
	db, err := strconv.Atoi(utils.GetEnvVar("REDIS_DB"))
	if err != nil {
		utils.ErrLog.Printf("could not parse REDIS_DB environment value: %v", err)
		return nil
	}

	expires, err := time.ParseDuration(utils.GetEnvVar("SECS_TO_EXPIRE"))
	if err != nil {
		utils.ErrLog.Printf("could not parse SECS_TO_EXPIRE environment value: %v", err)
		return nil
	}

	return NewRedisCache(address, db, expires)
}
