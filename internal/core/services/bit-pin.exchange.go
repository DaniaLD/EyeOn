package service

import (
	"context"
	"github.com/DaniaLD/EyeOn/internal/core/ports"
	"github.com/DaniaLD/EyeOn/internal/models"
)

type BitPinExchangeService struct {
	exchange ports.ExchangePort
}

func NewBitpinExchangeService(exchange ports.ExchangePort) *BitPinExchangeService {
	return &BitPinExchangeService{exchange: exchange}
}

func (s *BitPinExchangeService) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error) {
	return s.exchange.CreateOrder(ctx, req)
}

func (s *BitPinExchangeService) CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	return s.exchange.CancelOrder(ctx, req)
}

func (s *BitPinExchangeService) GetBalance(ctx context.Context) (*models.BalanceResponse, error) {
	return s.exchange.GetBalance(ctx)
}

func (s *BitPinExchangeService) GetOrderBook(ctx context.Context, req models.OrderBookRequest) (*models.OrderBookResponse, error) {
	return s.exchange.GetOrderBook(ctx, req)
}
