package service

import (
	"context"
	"github.com/DaniaLD/EyeOn/internal/core/ports"
	"github.com/DaniaLD/EyeOn/internal/models"
)

type ExchangeService struct {
	exchange ports.ExchangePort
}

func NewBitpinExchangeService(exchange ports.ExchangePort) *ExchangeService {
	return &ExchangeService{exchange: exchange}
}

func (s *ExchangeService) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error) {
	return s.exchange.CreateOrder(ctx, req)
}
