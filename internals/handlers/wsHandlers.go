package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/ctx"

	"github.com/gofiber/websocket/v2"
)

type Message struct {
	Text string `json:"text"`
}

func EnrollmentSessionHandler(c *websocket.Conn) {
	channel := "enroll"
	pubsub := cache.Client.Redis.Subscribe(ctx.Ctx, channel)

	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			var message Message

			message.Text = msg.Payload

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				return
			}

			if err := c.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				return
			} else {
				fmt.Println(jsonMessage)
			}
		}
	}()

	// for {
	// 	_, msg, err := c.ReadMessage()
	// 	if err != nil {
	// 		return
	// 	}

	// 	var message Message
	// 	err = json.Unmarshal(msg, &message)
	// 	if err != nil {
	// 		return
	// 	}

	// 	cache.Client.Redis.Publish(ctx.Ctx, channel, message.Text)
	// }

}
