package models

import "time"

type User struct {
	Id            int64     `gorm:"primary_key"`
	Name          string    `gorm:"name"`
	LastLoginTime time.Time `gorm:"last_login_time"`
	CreateTime    time.Time `gorm:"create_time"`
	UpdateTime    time.Time `gorm:"update_time"`
}
