package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Http struct {
	App  *fiber.App
	Conf *HttpConfig
}
type HttpConfig struct {
	Port string
}

func (h *Http) Setup() {
	h.App = fiber.New()
	h.App.Use(logger.New())
	h.App.Use(cors.New())
}
func (h *Http) Start() {
	h.App.Listen(fmt.Sprintf(":%s", h.Conf.Port))
}
func (h *Http) Shutdown() {
	h.App.Shutdown()
}
