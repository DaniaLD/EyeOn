package service

import (
	"context"
	"github.com/DaniaLD/EyeOn/internal/core/ports"
	"github.com/DaniaLD/EyeOn/internal/models"
)

type NobitexExchangeService struct {
	exchange ports.ExchangePort
}

func NewNobitexExchangeService(exchange ports.ExchangePort) *NobitexExchangeService {
	return &NobitexExchangeService{exchange: exchange}
}

func (s *NobitexExchangeService) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error) {
	return s.exchange.CreateOrder(ctx, req)
}

func (s *NobitexExchangeService) CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	return s.exchange.CancelOrder(ctx, req)
}

func (s *NobitexExchangeService) GetBalance(ctx context.Context) (*models.BalanceResponse, error) {
	return s.exchange.GetBalance(ctx)
}

func (s *NobitexExchangeService) GetOrderBook(ctx context.Context, req models.OrderBookRequest) (*models.OrderBookResponse, error) {
	return s.exchange.GetOrderBook(ctx, req)
}
