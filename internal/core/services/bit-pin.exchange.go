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

func (s *ExchangeService) CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	return s.exchange.CancelOrder(ctx, req)
}
