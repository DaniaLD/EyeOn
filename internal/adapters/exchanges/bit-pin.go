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
)

type BitpinExchange struct {
	token     string
	client    *http.Client
	baseURL   string
	apiKey    string
	secretKey string
}

func NewBitPinClient(apiKey, secretKey string) (ports.ExchangePort, error) {
	client := &BitpinExchange{
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    &http.Client{},
		baseURL:   "https://api.bitpin.ir",
	}

	return client, nil
}

func (b *BitpinExchange) authenticate(ctx context.Context) error {
	body := map[string]string{
		"api_key":    b.apiKey,
		"secret_key": b.secretKey,
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequestWithContext(ctx, "POST", b.baseURL+"/api/v1/usr/authenticate/", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("authentication failed: %s", resBody)
	}

	var result struct {
		Token string `json:"access"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	b.token = result.Token
	return nil
}

func (b *BitpinExchange) ensureToken(ctx context.Context) error {
	if b.token == "" {
		return b.authenticate(ctx)
	}
	return nil
}

func (b *BitpinExchange) authRequest(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+b.token)
	req.Header.Set("Content-Type", "application/json")
}

func (b *BitpinExchange) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (*models.OrderResponse, error) {
	if err := b.ensureToken(ctx); err != nil {
		return nil, err
	}

	body := map[string]interface{}{
		"symbol":      req.Symbol,
		"base_amount": req.Quantity,
		"price":       req.Price,
		"side":        string(req.Side),
		"type":        string(req.Type),
	}
	payload, _ := json.Marshal(body)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", b.baseURL+"/api/v1/odr/orders/", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	b.authRequest(httpReq)

	resp, err := b.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bitpin error: %s", resBody)
	}

	var result struct {
		ID     string `json:"id"`
		Status string `json:"state"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &models.OrderResponse{
		OrderID: result.ID,
		Symbol:  req.Symbol,
		Status:  result.Status,
	}, nil
}

func (b *BitpinExchange) CancelOrder(ctx context.Context, req models.CancelOrderRequest) (*models.CancelOrderResponse, error) {
	if err := b.ensureToken(ctx); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/api/v1/odr/orders/%s/", b.baseURL, req.OrderID), nil)
	if err != nil {
		return nil, err
	}
	b.authRequest(httpReq)

	resp, err := b.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bitpin error: %s", resBody)
	}

	return &models.CancelOrderResponse{Cancelled: true}, nil
}

func (b *BitpinExchange) GetBalance(ctx context.Context) (*models.BalanceResponse, error) {
	if err := b.ensureToken(ctx); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "GET", b.baseURL+"/api/v1/wlt/wallets/", nil)
	if err != nil {
		return nil, err
	}
	b.authRequest(httpReq)

	resp, err := b.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bitpin error: %s", bodyBytes)
	}

	var result []struct {
		Asset   string `json:"asset"`
		Balance string `json:"balance"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	assets := make(map[string]float64)
	for _, wallet := range result {
		balanceFloat, err := strconv.ParseFloat(wallet.Balance, 64)
		if err != nil {
			continue
		}
		assets[wallet.Asset] = balanceFloat
	}

	return &models.BalanceResponse{Assets: assets}, nil
}
