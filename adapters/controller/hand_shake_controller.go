package controller

import (
	"mhakheancode/api/ports/services"

	"github.com/gofiber/fiber/v2"
)

type HandshakeController struct {
	HandShakePort services.HandShakePort
}

func (h *HandshakeController) GetHandShake(c *fiber.Ctx) error {
	return c.SendString(h.HandShakePort.GetHandShake())
}
