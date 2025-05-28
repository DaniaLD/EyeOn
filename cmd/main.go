package main

import (
	"github.com/DaniaLD/EyeOn/internal/adapters/api"
	handler "github.com/DaniaLD/EyeOn/internal/adapters/api/handlers"
	"github.com/DaniaLD/EyeOn/internal/adapters/exchanges"
	service "github.com/DaniaLD/EyeOn/internal/core/services"
	"github.com/DaniaLD/EyeOn/pkg/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	configs.LoadConfigs()
	cfgs := configs.GetConfigs()
	bitpinApiKey := cfgs.BitPin.ApiKey
	bitpinSecretKey := cfgs.BitPin.SecretKey

	validator.New(validator.WithRequiredStructEnabled())

	if bitpinApiKey == "" || bitpinSecretKey == "" {
		log.Fatal("Missing BITPIN_API_KEY or BITPIN_SECRET_KEY env vars")
	}

	// Init services
	bitpinClient, err := exchanges.NewBitPinClient(bitpinApiKey, bitpinSecretKey)
	if err != nil {
		log.Fatalf("Unable to create bitpin client: %s", err.Error())
	}
	bitpinSvc := service.NewBitpinExchangeService(bitpinClient)
	bitpinHndlr := handler.NewBitpinHandler(bitpinSvc)

	engine := gin.Default()

	// Setup router
	r := api.NewRouter(engine, bitpinHndlr)
	r.Init()

	if err := engine.Run(":8080"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
