package redis

import (
	"message/config"
	"message/pkg/redis"
	"strconv"
)

var logger = config.Logger.WithField("mod", "redis")

func GetUsersInRoom(roomId int64) []int64 {
	return redis.ConvertSetToArray(redis.RDB.SMembers("room-" + strconv.Itoa(int(roomId))))
}

func EnterRoom(id int64, roomId int64) {
	redis.RDB.SAdd("room-"+strconv.Itoa(int(roomId)), id)
}

func LeaveRoom(id int64, roomId int64) {
	redis.RDB.SRem("room-"+strconv.Itoa(int(roomId)), id)
}

func IsInRoom(id int64, roomId int64) bool {
	return redis.RDB.SIsMember("room-"+strconv.Itoa(int(roomId)), id).Val()
}
