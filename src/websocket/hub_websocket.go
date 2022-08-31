package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"
)

type Hub struct {
	Clients     map[*Client]bool
	Broadcast   chan []byte
	Register    chan *Client
	UnRegister  chan *Client
	AddChatRoom chan *schemas.AddParticipantRequest
}

func NewHub() *Hub {
	return &Hub{
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan []byte),
		Register:    make(chan *Client),
		UnRegister:  make(chan *Client),
		AddChatRoom: make(chan *schemas.AddParticipantRequest),
	}
}

func (h *Hub) Run() {
	templateMessage := &models.Message{
		MessageId:     utils.GenerateId(),
		ChatRoomId:    "0",
		SenderId:      "0",
		StatusMessage: "",
		Type:          "noreply",
	}

	for {
		select {
		case client := <-h.Register:
			userId := client.UserId

			templateMessage.CreatedAt = time.Now().UnixMilli()
			templateMessage.Message = fmt.Sprintf("some one join room (ID: %s )", userId)
			msg, _ := json.Marshal(templateMessage)

			for client := range h.Clients {
				client.send <- msg
			}

			h.Clients[client] = true

		case client := <-h.UnRegister:
			userId := client.UserId
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}

			templateMessage.CreatedAt = time.Now().UnixMilli()
			templateMessage.Message = fmt.Sprintf("some one leave room (ID: %s )", userId)
			msg, _ := json.Marshal(templateMessage)

			for client := range h.Clients {
				client.send <- msg
			}

		case userMessage := <-h.Broadcast:
			var data map[string][]byte
			var message map[string]string
			json.Unmarshal(userMessage, &data)
			json.Unmarshal(data["message"], &message)

			roomId := message["chat_room_id"]

			for client := range h.Clients {
				if !client.ChatRoomsId[roomId] {
					continue
				}

				//prevent self receive the message
				if client.UserId == string(data["id"]) {
					continue
				}

				select {
				case client.send <- data["message"]:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}

		case participants := <-h.AddChatRoom:
			addingUserId := make(map[string]bool)

			for _, userId := range participants.UserIds {
				addingUserId[userId] = true
			}

			for client := range h.Clients {
				if addingUserId[client.UserId] {
					client.ChatRoomsId[participants.ChatRoomId] = true
				}
			}
		}

	}
}
