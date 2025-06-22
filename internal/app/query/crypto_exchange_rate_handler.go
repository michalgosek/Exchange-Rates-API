package query

import (
	"context"
	"exchange-rates-api/internal/core"
	"fmt"
)

type CryptoExchangeRateService interface {
	CalculateExchangeRate(ctx context.Context, from, to core.CurrencyCode, amount core.Decimal) (core.CalculatedExchangeRate, error)
}

type CryptoExchangeRateHandler struct {
	service CryptoExchangeRateService
}

func (c *CryptoExchangeRateHandler) Handle(ctx context.Context, query *CryptoExchangeRateQuery) (core.CalculatedExchangeRate, error) {
	dec, err := core.NewDecimal(query.Amount())
	if err != nil {
		return core.CalculatedExchangeRate{}, fmt.Errorf("failed to create decimal from float64: %w", err)
	}

	from, err := core.NewCryptoCurrencyCode(query.from)
	if err != nil {
		return core.CalculatedExchangeRate{}, fmt.Errorf("failed to create crypto currency code: %w", err)
	}

	to, err := core.NewCryptoCurrencyCode(query.to)
	if err != nil {
		return core.CalculatedExchangeRate{}, fmt.Errorf("failed to create crypto currency code: %w", err)
	}

	exchange, err := c.service.CalculateExchangeRate(ctx, from, to, dec)
	if err != nil {
		return core.CalculatedExchangeRate{}, fmt.Errorf("failed to calculcate exchange rate: %w", err)
	}

	return exchange, nil
}

func NewCryptoExchangeRateHandler(service CryptoExchangeRateService) *CryptoExchangeRateHandler {
	if service == nil {
		panic("crypto exchange rate service is required")
	}

	return &CryptoExchangeRateHandler{service: service}
}
