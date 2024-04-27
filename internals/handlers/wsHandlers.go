package handlers

import (
	"encoding/json"
	"sync"

	"github.com/AaronDennis07/electrum/internals/hub"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)
type Message struct {
	CourseId string `json:"course_id"`
}

var tempSeats = map[string]int{
	"course1": 10,
	"course2": 10,
	"course3": 10,
}
var mut = &sync.Mutex{}

func WsHandler(c *websocket.Conn) {
	defer func() {
		hub.Unregister <- c
		c.Close()
	}()

	hub.Register <- c
	jsonSeats , err := json.Marshal(tempSeats)
	if err!=nil{
		log.Error("could not marshal seats: ",err)
	}
	c.WriteMessage(websocket.TextMessage, jsonSeats)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("read:", err)
			}
			return
		}

		var message Message
		json.Unmarshal(msg, &message)

		//TODO: updatedb
		//TODO: updateredis



		//temp
		
		if _ , ok := tempSeats[message.CourseId]; ok{
			mut.Lock()
			tempSeats[message.CourseId]--
			mut.Unlock()

			err = c.WriteMessage(websocket.TextMessage, []byte("you are registered to the course: " + message.CourseId))
			if err!=nil{
			log.Error("could not confirm to client") //how to handle? remove from db or are they registered?
			return
		}
		}else{
			_ = c.WriteMessage(websocket.TextMessage, []byte("invalid course"))
		}
		

		jsonSeats , err := json.Marshal(tempSeats)
		if err!=nil{
			log.Error("could not marshal seats: ",err)
			break
		}
		hub.Broadcast <- string(jsonSeats)
		//temp

		//TODO: hub.Broadcast from redis
		//TODO: only register to one course
		
		




	}

}
