package services

import (
	"log"
	"net/http"

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

func BroadcastNotification(message string) {
	if SocketServer != nil {
		SocketServer.BroadcastToNamespace("/", "notification", message)
	}
}
