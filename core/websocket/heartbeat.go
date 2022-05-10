package websocket

import (
	"message/core/redis"
	"message/pkg/db"
	"time"
)

func heartbeat() {
	for {
		logger.Debugln("Current online:", len(manager.clients))
		for k, _ := range manager.clients {
			if !redis.IsOnline(k) {
				db.DB.Where("id = ?", k).Update("last_login_time", time.Now())
			}
			redis.Register(k)
		}
		<-time.After(time.Second * 20)
	}
}
