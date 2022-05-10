package db

import (
	"fmt"
	"message/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type BaseModel struct {
	ID int64 `gorm:"primary_key"`
}

func Initilize() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Global.DB_USERNAME, config.Global.DB_PASSWORD, config.Global.DB_HOST, config.Global.DB_PORT, config.Global.DB_DATABASE)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic(err)
	}
}
