package query

import (
	"context"
	"exchange-rates-api/internal/core"
	"fmt"
)

type GlobalRatesProvider interface {
	GetLatestExchangeRates(ctx context.Context, query *GlobalExchangeRatesQuery) ([]core.GlobalCurrencyExchangeRate, error)
}

type GlobalExchangeRateService interface {
	CalculateExchangeRates(ctx context.Context, rates ...core.GlobalCurrencyExchangeRate) ([]core.CalculatedExchangeRate, error)
}

type GlobalExchangeRatesHandler struct {
	provider GlobalRatesProvider
	service  GlobalExchangeRateService
}

func (g *GlobalExchangeRatesHandler) Handle(ctx context.Context, query *GlobalExchangeRatesQuery) ([]core.CalculatedExchangeRate, error) {
	rates, err := g.provider.GetLatestExchangeRates(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest exchange rate: %w", err)
	}

	exchanges, err := g.service.CalculateExchangeRates(ctx, rates...)
	if err != nil {
		return nil, fmt.Errorf("failed to calculcate exchange rates: %w", err)
	}

	return exchanges, nil
}

func NewGlobalExchangeRatesHandler(provider GlobalRatesProvider, service GlobalExchangeRateService) *GlobalExchangeRatesHandler {
	if provider == nil {
		panic("global rates provider is required")
	}
	if service == nil {
		panic("global exchange rate service is required")
	}

	return &GlobalExchangeRatesHandler{service: service, provider: provider}
}
