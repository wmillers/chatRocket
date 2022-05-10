package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Viper  *viper.Viper
	Global Config
	Logger = logrus.New()
)

type Config struct {
	APP_NAME      string
	APP_DEV_MODE  bool
	APP_YM        string
	APP_NODE      string
	APP_GOROUTINE int

	WEBSOCKET_PORT string
	TCP_PORT       string
	GRPC_PORT      string
	SWARM_PORT     string
	LOG_ADDR       string

	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	DB_LOC      string

	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       int

	RABBITMQ_HOST     string
	RABBITMQ_PORT     string
	RABBITMQ_USER     string
	RABBITMQ_PASSWORD string

	JWT_SECRET     string
	JWT_EXPIRES    int
	BASE64_ENCRYPT string

	APP_CLUSTER_MODEL bool
}

func init() {
	Viper = viper.New()
	Viper.SetConfigName("config")
	Viper.SetConfigType("toml")
	Viper.AddConfigPath(".")

	if err := Viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := Viper.Unmarshal(&Global); err != nil {
		panic(err)
	}
	logrus.WithField("mod", "conf").Println(Global)
}

func Initilize() {}
