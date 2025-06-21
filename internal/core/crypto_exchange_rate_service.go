package core

import (
	"context"
	"fmt"
)

type CryptoExchangeArithmeticService interface {
	DivWithPrecision(ExchangeRate, ExchangeRate, DecimalPrecision) (ExchangeRate, error)
	MulWithPrecision(ExchangeRate, ExchangeRate, DecimalPrecision) (ExchangeRate, error)
}

type CryptoExchangeRateService struct {
	service CryptoExchangeArithmeticService
	table   CryptoExchangeRateTable
}

func (c *CryptoExchangeRateService) CalculateExchangeRate(ctx context.Context, from, to CurrencyCode, amount Decimal) (CalculatedExchangeRate, error) {
	tableEntryFrom, err := c.table.GetExchangeRate(from)
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	tableEntryTo, err := c.table.GetExchangeRate(to)
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	quantity, err := NewExchangeRate(amount.String())
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to create exchange rate: %w", err)
	}

	mul, err := c.service.MulWithPrecision(tableEntryFrom.Rate(), quantity, tableEntryTo.DecimalPrecision())
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to multiply rates (%s x %s): %w", tableEntryFrom.Rate(), quantity, err)
	}

	// from the first rate (EUR) to the second rate (USD) i.e. EUR -> USD
	div, err := c.service.DivWithPrecision(mul, tableEntryTo.Rate(), tableEntryTo.DecimalPrecision())
	if err != nil {
		return CalculatedExchangeRate{}, fmt.Errorf("failed to div rates (%s / %s): %w", mul.String(), tableEntryFrom.Rate().String(), err)
	}

	return NewCalculatedExchangeRate(from, to, div), nil
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
