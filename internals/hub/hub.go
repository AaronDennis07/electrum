package hub

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client)
var Register = make(chan *websocket.Conn)
var Unregister = make(chan *websocket.Conn)
var Broadcast = make(chan string)

func RunHub() {
	for {
		select {
			case conn := <-Register:
				clients[conn] = &client{}
				log.Info("Client registered: ", conn.IP() , " n: ", len(clients))

			case conn := <-Unregister:
				delete(clients, conn)
				log.Info("Client unregistered: ", conn.IP() , " n: ", len(clients))
			
			case message := <-Broadcast:
				log.Info("message received:", message)
			// Send the message to all clients
				for conn, c := range clients {
					go func(conn *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
						c.mu.Lock()
						defer c.mu.Unlock()
						if c.isClosing {
							return
						}
						if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
							c.isClosing = true
							log.Error("write error:", err)

							conn.WriteMessage(websocket.CloseMessage, []byte{})
							conn.Close()
							Unregister <- conn
						}
					}(conn, c)
				}

		}
	}
}