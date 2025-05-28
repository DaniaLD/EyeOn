package models

type Side string
type OrderType string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"

	OrderTypeLimit     OrderType = "limit"
	OrderTypeMarket    OrderType = "market"
	OrderTypeOCO       OrderType = "oco"
	OrderTypeStopLimit OrderType = "stop_limit"
)

type CreateOrderRequest struct {
	Symbol   string    `form:"symbol" binding:"required"`
	Side     Side      `form:"side" binding:"required,oneof=buy sell"`
	Type     OrderType `form:"type" binding:"required,oneof=limit market oco stop_limit"`
	Price    float64   `form:"price" binding:"required"`
	Quantity float64   `form:"quantity" binding:"required"`
}

type OrderResponse struct {
	OrderID string
	Symbol  string
	Status  string
}

type CancelOrderRequest struct {
	OrderID string `uri:"id" binding:"required"`
}

type CancelOrderResponse struct {
	Cancelled bool
}

type BalanceResponse struct {
	Assets map[string]float64
}
