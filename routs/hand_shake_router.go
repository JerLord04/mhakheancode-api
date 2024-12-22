package routs

import (
	"mhakheancode/api/adapters/controller"
	"mhakheancode/api/domain"

	"github.com/gofiber/fiber/v2"
)

func NewHandShakeRouter(group fiber.Router) {
	handshakeController := &controller.HandshakeController{
		HandShakePort: domain.NewHandShakeDomain(),
	}
	group.Get("/handShake", handshakeController.GetHandShake)
}
