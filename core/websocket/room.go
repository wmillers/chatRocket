package websocket

import "message/core/redis"

func roomSend(sender *Client, message Message) {
	users := redis.GetUsersInRoom(message.RoomId)
	logger.WithField("room", message.RoomId).WithField("sender", sender.id).Infoln("Message", message.Message)
	for _, user := range users {
		if user != message.ID && manager.clients[user] != nil {
			manager.clients[user].sendMsg(message.Message, sender, message.RoomId)
		}
	}
}

func broadcast(message []byte) {
	for _, v := range manager.clients {
		v.sendMsg(string(message), broadcaster, broadcaster.roomId)
	}
}
