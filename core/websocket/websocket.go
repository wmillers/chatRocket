package websocket

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"message/config"
	"message/core/redis"
	"message/models"
	"message/pkg/db"
)

var (
	logger      *logrus.Entry
	broadcaster = &Client{name: "broadcast"}
	manager     = &ClientManager{
		clients:    make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ClientManager struct {
	clients    map[int64]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	message    chan *Message
}

type Client struct {
	socket *websocket.Conn
	id     int64
	roomId int64
	name   string
}

type Message struct {
	Name    string `json:"name"`
	ID      int64  `json:"id"`
	RoomId  int64  `json:"roomId"`
	Message string `json:"message"`
}

func (manager *ClientManager) router() {
	for {
		select {
		case client := <-manager.register:
			logger.Infoln("New client registered", client.name, client.id)
			manager.clients[client.id] = client
			redis.Register(client.id)
			client.sendMsg("Welcome to the room!", broadcaster, client.id)
		case client := <-manager.unregister:
			logger.Infoln("Client unregistered", client.id)
			delete(manager.clients, client.id)
			redis.Unregister(client.id)
			client.sendMsg("You have been disconnected", broadcaster, broadcaster.roomId)
		case message := <-manager.broadcast:
			logger.Infoln("Broadcast message", message)
			broadcast(message)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Debugln("New connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Panicln(err)
		return
	}

	client := GenerateRandomClient(conn, 1)
	manager.register <- client
	go clientListener(client)
}

func clientListener(client *Client) {
	for {
		_, message, err := client.socket.ReadMessage()
		if err != nil {
			logger.Debugln(err)
			manager.unregister <- client
			client.socket.Close()
			return
		}
		logger.Debugln("Received message", string(message), "from", client.id)
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			logger.Errorln(err)
			continue
		}
		if msg.RoomId > 0 {
			client.roomId = msg.RoomId
			redis.EnterRoom(client.id, client.roomId)
			roomSend(client, msg)
		}
	}
}

func WebsocketStart() {
	logger = config.Logger.WithField("mod", "ws")
	logger.WithField("port", config.Global.WEBSOCKET_PORT).Println("Starting websocket server...")

	go manager.router()
	go heartbeat()

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+config.Global.WEBSOCKET_PORT, nil)
	if err != nil {
		logger.Panicln(err)
	}
}

func (client *Client) sendMsg(msg string, sender *Client, roomId int64) {
	client.socket.WriteJSON(Message{sender.name, sender.id, roomId, msg})
}

func GenerateRandomClient(conn *websocket.Conn, testRoom int64) *Client {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10000000)
	emoji := [][]int{
		// Emoticons icons
		{128513, 128591},
		// Transport and map symbols
		{128640, 128704},
	}
	i := r % 2
	name := string(r%(emoji[i][1]-emoji[i][0]) + emoji[i][0])
	redis.EnterRoom(int64(r), testRoom)
	db.DB.Create(&models.User{
		Id:            int64(r),
		Name:          strconv.Itoa(r),
		LastLoginTime: time.Now(),
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	})
	return &Client{
		socket: conn,
		id:     int64(r),
		name:   name,
		roomId: testRoom,
	}
}
