package main

import (
	"exchange-rates-api/internal/adapters"
	"exchange-rates-api/internal/app"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"exchange-rates-api/internal/infrastructure/config"
	"exchange-rates-api/internal/infrastructure/math"
	"exchange-rates-api/internal/infrastructure/server"
	"flag"
	"log"
)

//go:generate go tool oapi-codegen -generate types -o ../internal/ports/openapi_types.go -package ports ../api/openapi/exchange_rates.yaml
//go:generate go tool oapi-codegen -generate gin-server -o ../internal/ports/openapi_api.gen.go -package ports ../api/openapi/exchange_rates.yaml
func main() {
	// Configuration workflow initialization:
	path := flag.String("config", "../config.yaml", "Path to the HTTP server config file.")
	flag.Parse()
	cfg, err := config.LoadConfig(*path)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(cfg.String())

	// Dependencies workflow initialization:
	app := &app.Application{
		Queries: app.Queries{
			GlobalExchangeRatesHandler: query.NewGlobalExchangeRatesHandler(
				adapters.NewOpenExchangeRatesHTTP(&adapters.OpenExchangeRatesHTTPConfig{
					AppID:   cfg.OpenExchangeRatesAPI.AppID,
					BaseURL: cfg.OpenExchangeRatesAPI.BaseURL,
				}),
				core.NewGlobalExchangeRateService(math.NewCurrencyExchangeArithmeticService(math.CurrencyExchangeRatePrecision)),
			),
			CryptoExchangeRateHandler: query.NewCryptoExchangeRateHandler(
				core.NewCryptoExchangeRateService(
					math.NewCurrencyExchangeArithmeticService(math.DefaultExchangeRatePrecision),
					core.NewDefaultCryptoExchangeRateTable(),
				),
			),
		},
	}

	// HTTP Server workflow initialization:
	done := server.RunHTTPWithGracefulShutdown(cfg, app)
	<-done
}
