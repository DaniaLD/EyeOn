package handler

import (
	"context"
	service "github.com/DaniaLD/EyeOn/internal/core/services"
	"github.com/DaniaLD/EyeOn/internal/models"
	dtovalidator "github.com/DaniaLD/EyeOn/pkg/dto-validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BitpinHandler struct {
	svc *service.BitPinExchangeService
}

func NewBitpinHandler(svc *service.BitPinExchangeService) *BitpinHandler {
	return &BitpinHandler{svc: svc}
}

func (h *BitpinHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if !dtovalidator.BindBodyAndValidate(c, &req) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := h.svc.CreateOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BitpinHandler) CancelOrder(c *gin.Context) {
	var req models.CancelOrderRequest
	if !dtovalidator.BindUriAndValidate(c, &req) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := h.svc.CancelOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BitpinHandler) GetBalance(c *gin.Context) {
	resp, err := h.svc.GetBalance(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *BitpinHandler) GetOrderBook(c *gin.Context) {
	var req models.OrderBookRequest
	if !dtovalidator.BindUriAndValidate(c, &req) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := h.svc.GetOrderBook(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
