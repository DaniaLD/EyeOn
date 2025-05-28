package api

import (
	handler "github.com/DaniaLD/EyeOn/internal/adapters/api/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine        *gin.Engine
	bitpinHandler *handler.BitpinHandler
}

func NewRouter(
	engine *gin.Engine,
	bitpinHandler *handler.BitpinHandler) Router {
	return Router{
		engine:        engine,
		bitpinHandler: bitpinHandler,
	}
}

func (r Router) Init() {
	v1 := r.engine.Group("/api/v1")

	// Bitpin routes
	v1.POST("/bitpin/order", r.bitpinHandler.CreateOrder)
	v1.DELETE("/bitpin/order/:id", r.bitpinHandler.CancelOrder)
	v1.GET("bitpin/wallet", r.bitpinHandler.GetBalance)
	v1.GET("bitpin/order-book/:symbol", r.bitpinHandler.GetOrderBook)
}
