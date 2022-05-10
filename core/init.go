package core

import (
	"message/config"
	"message/pkg/db"
	"message/pkg/redis"
)

func Initilize() {
	logger := config.Logger.WithField("mod", "init")
	logger.Println("Initializing...")

	db.Initilize()
	redis.Initilize()
}
