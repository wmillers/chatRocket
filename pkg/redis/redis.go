package redis

import (
	"message/config"
	"strconv"

	"github.com/go-redis/redis"
)

var RDB *redis.Client

func Initilize() {
	logger := config.Logger.WithField("mod", "init")
	RDB = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     config.Global.REDIS_HOST + ":" + config.Global.REDIS_PORT,
		Password: config.Global.REDIS_PASSWORD,
		DB:       config.Global.REDIS_DB,
		PoolSize: 15,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		logger.Fatal("Redis connection failed")
	}
}

func ConvertSetToArray(set *redis.StringSliceCmd) []int64 {
	var ids []int64
	for _, user := range set.Val() {
		id, _ := strconv.ParseInt(user, 10, 64)
		ids = append(ids, id)
	}
	return ids
}
