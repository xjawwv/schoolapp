package services

import (
	"log"
	"net/http"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var SocketServer *socketio.Server

func InitSocketServer() *socketio.Server {
	opts := &engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	}
	server := socketio.NewServer(opts)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("Socket.IO client connected:", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Socket.IO client disconnected:", s.ID(), "reason:", reason)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Socket.IO error:", e)
	})

	SocketServer = server
	return server
}

func BroadcastNotification(title string, description string) {
	notification := models.Notification{
		Title:       title,
		Description: description,
		IsRead:      false,
	}
	if config.DB != nil {
		config.DB.Create(&notification)
	}

	if SocketServer != nil {
		SocketServer.BroadcastToNamespace("/", "notification", map[string]interface{}{
			"id":          notification.ID,
			"title":       notification.Title,
			"description": notification.Description,
			"created_at":  notification.CreatedAt,
		})
	}
}

func BroadcastPopup(title string, photo string) {
	popup := models.Popup{
		Title: title,
		Photo: photo,
	}
	if config.DB != nil {
		config.DB.Create(&popup)
	}

	if SocketServer != nil {
		SocketServer.BroadcastToNamespace("/", "popup", map[string]interface{}{
			"id":         popup.ID,
			"title":      popup.Title,
			"photo":      popup.Photo,
			"created_at": popup.CreatedAt,
		})
	}
}
