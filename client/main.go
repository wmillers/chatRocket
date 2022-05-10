package main

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Message struct {
	Name    string `json:"name"`
	ID      int64  `json:"id"`
	RoomId  int64  `json:"roomId"`
	Message string `json:"message"`
}

func WSClient() {
	logger := logrus.New().WithField("mod", "wsc")
	conn, err := GetWSConn()
	if err != nil {
		panic(err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			panic(err)
		}
		if msg.ID == 0 && msg.Name == "broadcast" {
			logger = logger.WithField("id", msg.RoomId)
		}
		logger.Infoln(msg)
	}
}

func WSClientSend() {
	conn, err := GetWSConn()
	if err != nil {
		panic(err)
	}
	msg := Message{
		Name:    "broadcast",
		ID:      0,
		RoomId:  1,
		Message: "hello",
	}
	err = conn.WriteJSON(msg)
	if err != nil {
		panic(err)
	}
}

func GetWSConn() (*websocket.Conn, error) {
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial("ws://localhost:9123", nil)
	return conn, err
}

func main() {
	for i := 0; i < 2; i++ {
		go WSClient()
	}
	<-time.After(time.Second * 1)
	WSClientSend()
	<-time.After(time.Hour)
}
