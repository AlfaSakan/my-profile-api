package websocket

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserConn struct {
	UserId  string `json:"user_id"`
	Address string `json:"address"`
	EnterAt int64  `json:"enter_at"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	UserConn

	ChatRoomsId map[string]bool
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.UnRegister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		data := map[string][]byte{
			"message": message,
			"id":      []byte(c.UserId),
		}
		userMessage, _ := json.Marshal(data)
		c.hub.Broadcast <- userMessage
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(c *gin.Context, hub *Hub, chatRooms *[]schemas.ChatRoomWithPartisipants) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	userId := c.Param("UserId")

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), ChatRoomsId: make(map[string]bool)}
	client.hub.Register <- client
	client.UserId = userId
	client.Address = conn.RemoteAddr().String()
	client.EnterAt = time.Now().UnixMilli()

	for _, room := range *chatRooms {
		client.ChatRoomsId[room.ChatRoomId] = true
	}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	welcomeMessage := &models.Message{
		MessageId:     utils.GenerateId(),
		CreatedAt:     client.EnterAt,
		ChatRoomId:    "0",
		SenderId:      "0",
		StatusMessage: "",
		Type:          "noreply",
		Message:       "Welcome",
	}

	ArrByteMessage, _ := json.Marshal(welcomeMessage)

	client.send <- ArrByteMessage
}
