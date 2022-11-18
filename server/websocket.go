package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type WebsocketConfig struct {
}
type Websocket struct {
	App  *fiber.App
	Conf *WebsocketConfig
}

func (w *Websocket) Setup() {
	w.App.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	w.App.Get("/wss", websocket.New(w.wss_handler))

}
func (w *Websocket) Start() {
}
func (w *Websocket) Shutdown() {
}

func (w *Websocket) wss_handler(c *websocket.Conn) {
	for {
		messageType, message, err := c.ReadMessage()
		_ = messageType

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("read error:", err)
			}
			c.Close()
			break
		}

		fmt.Println("read:", message)

		if err := c.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
