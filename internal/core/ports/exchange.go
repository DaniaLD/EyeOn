package ports

import (
	"context"
	"github.com/DaniaLD/EyeOn/internal/models"
)

type ExchangePort interface {
	CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error)
	CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error)
}
