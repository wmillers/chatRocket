package redis

import (
	"message/pkg/redis"
	"strconv"
	"time"
)

func Register(id int64) {
	redis.RDB.Set("user-"+strconv.Itoa(int(id)), time.Now().Format("2006-01-02 15:04:05"), 30*time.Second)
}

func Unregister(id int64) {
	redis.RDB.Del("user-" + strconv.Itoa(int(id)))
}

func IsOnline(id int64) bool {
	return redis.RDB.Get("user-"+strconv.Itoa(int(id))).Val() != ""
}
