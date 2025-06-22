package math

import (
	"exchange-rates-api/internal/core"
	"fmt"

	"github.com/cockroachdb/apd/v3"
)

const (
	DefaultExchangeRatePrecision  = 50
	CurrencyExchangeRatePrecision = 5
)

type CurrencyExchangeArithmeticService struct {
	ctx  *apd.Context
	prec uint32
}

func (c *CurrencyExchangeArithmeticService) CalculateCrossRate(first, second core.ExchangeRate, amount core.Decimal) (core.ExchangeRate, error) {
	rate1, _, err := apd.NewFromString(first.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", first.String(), err)
	}

	rate2, _, err := apd.NewFromString(second.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", second.String(), err)
	}

	amt, _, err := apd.NewFromString(amount.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", amount.String(), err)
	}

	// amount × rate1
	multiplied := apd.New(0, 0)
	_, err = c.ctx.Mul(multiplied, amt, rate1)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to calculate product (%s x %s): %w", amt.String(), first.String(), err)
	}

	// (amount × rate1) / rate2
	result := apd.New(0, 0)
	_, err = c.ctx.Quo(result, multiplied, rate2)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to calculate quotient (%s / %s): %w", rate2.String(), multiplied.String(), err)
	}

	return core.NewExchangeRate(result.Text('f'))
}

func (c *CurrencyExchangeArithmeticService) CalculateCrossRateWithPrecission(first, second core.ExchangeRate, amount core.Decimal, prec core.DecimalPrecision) (core.ExchangeRate, error) {
	ex, err := c.CalculateCrossRate(first, second, amount)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create exchange cross rate: %w", err)
	}

	rate, _, err := apd.NewFromString(ex.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", ex.String(), err)
	}

	// Apply quantization to match 'to' decimal precision
	quant := apd.New(1, -int32(prec.Value()))
	final := apd.New(0, 0)
	_, err = c.ctx.Quantize(final, rate, quant.Exponent)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to quantize calculated quotient (%s): %w", quant.String(), err)
	}

	return core.NewExchangeRate(final.Text('f'))
}

func NewCurrencyExchangeArithmeticService(prec uint32) *CurrencyExchangeArithmeticService {
	if prec == 0 {
		panic("invalid precision: must be a number greater than 0")
	}

	return &CurrencyExchangeArithmeticService{prec: prec, ctx: apd.BaseContext.WithPrecision(prec)}
}
