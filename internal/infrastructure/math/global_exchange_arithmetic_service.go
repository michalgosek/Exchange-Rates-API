package math

import (
	"exchange-rates-api/internal/core"
	"fmt"

	"github.com/cockroachdb/apd/v3"
)

const (
	IEE754Precision              = 34
	DefaultCurrencyRatePrecision = 5
)

type GlobalExchangeArithmeticService struct {
	ctx  *apd.Context
	prec uint32
}

func (g *GlobalExchangeArithmeticService) MulWithPrecision(first, second core.ExchangeRate, prec core.DecimalPrecision) (core.ExchangeRate, error) {
	num1, _, err := apd.NewFromString(first.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", first.String(), err)
	}

	num2, _, err := apd.NewFromString(second.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", second.String(), err)
	}

	mul := apd.New(0, 0)
	_, err = g.ctx.Mul(mul, num1, num2)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to calculate product (%s x %s): %w", first.String(), second.String(), err)
	}

	quant := apd.New(1, -int32(prec.Value()))
	res := apd.New(0, 0)
	_, err = g.ctx.Quantize(res, mul, quant.Exponent)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to quantize calculated quotient (%s): %w", mul.String(), err)
	}

	return core.NewExchangeRate(res.Text('f'))
}

func (g *GlobalExchangeArithmeticService) DivWithPrecision(first, second core.ExchangeRate, prec core.DecimalPrecision) (core.ExchangeRate, error) {
	ex, err := g.Div(first, second)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to calculate quotient (%s / %s): %w", first.String(), second.String(), err)
	}

	quo, _, err := apd.NewFromString(ex.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", second.String(), err)
	}

	quant := apd.New(1, -int32(prec.Value()))
	res := apd.New(0, 0)
	_, err = g.ctx.Quantize(res, quo, quant.Exponent)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to quantize calculated quotient (%s): %w", quo.String(), err)
	}

	return core.NewExchangeRate(res.Text('f'))
}

func (g *GlobalExchangeArithmeticService) Div(first, second core.ExchangeRate) (core.ExchangeRate, error) {
	num1, _, err := apd.NewFromString(first.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", first.String(), err)
	}

	num2, _, err := apd.NewFromString(second.String())
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to create APD dec from str %s: %w", second.String(), err)
	}

	res := apd.New(0, 0)
	_, err = g.ctx.Quo(res, num1, num2)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to calculate quotient (%s / %s): %w", first.String(), second.String(), err)
	}

	return core.NewExchangeRate(res.Text('f'))
}

func NewGlobalExchangeArithmeticService(prec uint32) *GlobalExchangeArithmeticService {
	if prec == 0 {
		panic("invalid precision: must be a number greater than 0")
	}

	return &GlobalExchangeArithmeticService{prec: prec, ctx: apd.BaseContext.WithPrecision(prec)}
}
