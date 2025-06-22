package core

import (
	"context"
	"fmt"
)

type CryptoExchangeArithmeticService interface {
	CalculateCrossRateWithPrecission(first, second ExchangeRate, amount Decimal, to DecimalPrecision) (ExchangeRate, error)
}

type CryptoExchangeRateService struct {
	service CryptoExchangeArithmeticService
	table   CryptoExchangeRateTable
}

func (c *CryptoExchangeRateService) CalculateExchangeRate(ctx context.Context, from, to CurrencyCode, amount Decimal) (CalculatedExchangeRate, error) {
	first, err := c.table.GetExchangeRate(from)
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	second, err := c.table.GetExchangeRate(to)
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	rate, err := c.service.CalculateCrossRateWithPrecission(first.Rate(), second.Rate(), amount, second.DecimalPrecision())
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to calculate cross rate: %w", err)
	}

	return NewCalculatedExchangeRate(from, to, rate), nil
}

func NewCryptoExchangeRateService(service CryptoExchangeArithmeticService, table CryptoExchangeRateTable) *CryptoExchangeRateService {
	if service == nil {
		panic("crypto exchange arithmetic service is required")
	}
	if table == nil {
		panic("crypto exchange rate table is required")
	}

	return &CryptoExchangeRateService{service: service, table: table}
}
