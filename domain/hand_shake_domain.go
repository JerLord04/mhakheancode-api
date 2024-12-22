package domain

import "mhakheancode/api/ports/services"

type handShakeDomain struct {
}

func (h *handShakeDomain) GetHandShake() string {
	return "Handshake with fiber api. 🙋"
}

func NewHandShakeDomain() services.HandShakePort {
	return &handShakeDomain{}
}
