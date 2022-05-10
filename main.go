package main

import (
	"flag"
	"message/config"
	"message/core"
	"message/core/websocket"

	"github.com/sirupsen/logrus"
)

var ServerType string

func init() {
	flag.StringVar(&ServerType, "serve", "", "serve type")
	flag.Parse()
	config.Initilize()
	if config.Global.APP_DEV_MODE {
		config.Logger.SetLevel(logrus.DebugLevel)
	}
	core.Initilize()
}

func main() {
	websocket.WebsocketStart()
}
