package testabilities

import (
	"exchange-rates-api/internal/core"
	"fmt"
	"math"
	"strconv"
	"testing"
)

const DefaultRoundPrecision = 5

type TestGlobalExchangeArithmeticServiceFloat64 struct {
	t     *testing.T
	round float64
}

func (t TestGlobalExchangeArithmeticServiceFloat64) DivWithPrecision(first, second core.ExchangeRate, prec core.DecimalPrecision) (core.ExchangeRate, error) {
	t.t.Helper()

	x, err := strconv.ParseFloat(first.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", first.String(), err)
	}
	y, err := strconv.ParseFloat(second.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", second.String(), err)
	}

	div := x / y
	factor := math.Pow(10, float64(prec.Value()))
	rounded := math.Round(div*factor) / factor

	return core.NewExchangeRate(strconv.FormatFloat(rounded, 'g', -1, 64))
}

func (t TestGlobalExchangeArithmeticServiceFloat64) MulWithPrecision(first, second core.ExchangeRate, prec core.DecimalPrecision) (core.ExchangeRate, error) {
	t.t.Helper()

	x, err := strconv.ParseFloat(first.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", first.String(), err)
	}
	y, err := strconv.ParseFloat(second.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", second.String(), err)
	}

	prod := x * y
	factor := math.Pow(10, float64(prec.Value()))
	rounded := math.Round(prod*factor) / factor

	return core.NewExchangeRate(strconv.FormatFloat(rounded, 'g', -1, 64))
}

func (t TestGlobalExchangeArithmeticServiceFloat64) Div(first core.ExchangeRate, second core.ExchangeRate) (core.ExchangeRate, error) {
	t.t.Helper()

	x, err := strconv.ParseFloat(first.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", first.String(), err)
	}
	y, err := strconv.ParseFloat(second.String(), 64)
	if err != nil {
		return core.ExchangeRate{}, fmt.Errorf("failed to parse str %s to float64: %w", second.String(), err)
	}

	div := x / y
	factor := math.Pow(10, t.round)
	rounded := math.Round(div*factor) / factor

	return core.NewExchangeRate(strconv.FormatFloat(rounded, 'g', -1, 64))
}

func NewTestGlobalExchangeArithmeticServiceFloat64(t *testing.T, round float64) *TestGlobalExchangeArithmeticServiceFloat64 {
	if t == nil {
		panic("testing parameter is required")
	}
	if round <= 0 {
		panic("invalid round: must be a non-negative number greater than zero")
	}

	return &TestGlobalExchangeArithmeticServiceFloat64{t: t, round: round}
}
