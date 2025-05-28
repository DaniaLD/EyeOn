package exchanges

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DaniaLD/EyeOn/internal/core/ports"
	"github.com/DaniaLD/EyeOn/internal/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type NobitexExchange struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewNobitexClient(apiKey string) (ports.ExchangePort, error) {
	client := &NobitexExchange{
		apiKey:  apiKey,
		client:  &http.Client{},
		baseURL: "https://api.nobitex.ir",
	}

	return client, nil
}

func (n *NobitexExchange) authRequest(req *http.Request) {
	req.Header.Set("Authorization", "Token "+n.apiKey)
}

func (n *NobitexExchange) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error) {
	symbol := strings.Split(req.Symbol, "_")
	body := map[string]interface{}{
		"srcCurrency": strings.ToLower(symbol[0]),
		"dstCurrency": strings.ToLower(symbol[1]),
		"amount":      req.Quantity,
		"price":       req.Price,
		"type":        string(req.Side),
		"execution":   string(req.Type),
	}
	payload, _ := json.Marshal(body)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", n.baseURL+"/market/orders/add", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	n.authRequest(httpReq)

	resp, err := n.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("nobitex error: %s", resBody)
	}

	var result struct {
		Order struct {
			ID     int    `json:"id"`
			Status string `json:"state"`
		} `json:"order"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	orderId := strconv.Itoa(result.Order.ID)
	return &models.OrderResponse{
		OrderID: orderId,
		Symbol:  req.Symbol,
		Status:  result.Order.Status,
	}, nil
}

func (n *NobitexExchange) CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	body := map[string]interface{}{
		"order":  req.OrderID,
		"status": "canceled",
	}
	payload, _ := json.Marshal(body)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/market/orders/update-status", n.baseURL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	n.authRequest(httpReq)

	resp, err := n.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("nobitex error: %s", resBody)
	}

	return &models.CancelOrderResponse{Cancelled: true}, nil
}

func (n *NobitexExchange) GetBalance(ctx context.Context) (*models.BalanceResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, "GET", n.baseURL+"/users/wallets/list", nil)
	if err != nil {
		return nil, err
	}
	n.authRequest(httpReq)

	resp, err := n.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("nobitex error: %s", bodyBytes)
	}

	var result struct {
		Wallets []struct {
			Balance  string `json:"balance"`
			Currency string `json:"currency"`
		} `json:"wallets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	assets := make(map[string]float64)
	for _, wallet := range result.Wallets {
		balanceFloat, err := strconv.ParseFloat(wallet.Balance, 64)

		if err != nil {
			continue
		}
		if balanceFloat != 0 {
			assets[wallet.Currency] = balanceFloat
		}
	}

	return &models.BalanceResponse{Assets: assets}, nil
}

func (n *NobitexExchange) GetOrderBook(ctx context.Context, req models.OrderBookRequest) (*models.OrderBookResponse, error) {
	symbol := strings.Replace(req.Symbol, "_", "", -1)
	httpReq, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/v2/trades/%s", n.baseURL, symbol), nil)
	if err != nil {
		return nil, err
	}

	resp, err := n.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("nobitex error: %s", resBody)
	}

	var result struct {
		Trades []struct {
			Price    string `json:"price"`
			Quantity string `json:"volume"`
			Type     string `json:"type"`
		} `json:"trades"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	bids := make([]models.OrderBookEntry, 0)
	asks := make([]models.OrderBookEntry, 0)
	for _, order := range result.Trades {
		if order.Type == "sell" {
			bids = append(bids, models.OrderBookEntry{
				Price:    order.Price,
				Quantity: order.Quantity,
			})
		} else {
			asks = append(asks, models.OrderBookEntry{
				Price:    order.Price,
				Quantity: order.Quantity,
			})
		}
	}

	return &models.OrderBookResponse{
		Bids: bids,
		Asks: asks,
	}, nil
}
