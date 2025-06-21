package main

import (
	"context"
	"exchange-rates-api/internal/adapters"
	"exchange-rates-api/internal/app"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"exchange-rates-api/internal/infrastructure/config"
	"exchange-rates-api/internal/infrastructure/math"
	"exchange-rates-api/internal/ports"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	log.Println("CONFIG LOADED")

	// Dependencies workflow initialization:
	arithmeticService := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)
	globalExchangeService := core.NewGlobalExchangeRateService(math.NewGlobalExchangeArithmeticService(math.DefaultCurrencyRatePrecision))
	cryptoExchangeService := core.NewCryptoExchangeRateService(arithmeticService, core.NewDefaultCryptoExchangeRateTable())
	globalRatesService := adapters.NewOpenExchangeRatesHTTP(&adapters.OpenExchangeRatesHTTPConfig{
		AppID:   cfg.OpenExchangeRatesAPI.AppID,
		BaseURL: cfg.OpenExchangeRatesAPI.BaseURL,
	})

	// HTTP Server workflow initialization:
	srv := http.Server{
		Addr: cfg.Server.SocketAddr(),
		Handler: ports.NewHTTP(&app.Application{
			Queries: app.Queries{
				GlobalExchangeRatesHandler: query.NewGlobalExchangeRatesHandler(globalRatesService, globalExchangeService),
				CryptoExchangeRateHandler:  query.NewCryptoExchangeRateHandler(cryptoExchangeService),
			},
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		log.Print("HTTP graceful shutdown")
		close(done)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-done
}
