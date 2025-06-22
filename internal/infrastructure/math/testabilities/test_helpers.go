package testabilities

import (
	"exchange-rates-api/internal/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewExchangeRateTestHelper(t *testing.T, s string) core.ExchangeRate {
	v, err := core.NewExchangeRate(s)
	require.NoError(t, err)
	require.NotZero(t, v)
	return v
}

func NewDecimalPrecisionTestHelper(t *testing.T, u uint32) core.DecimalPrecision {
	v, err := core.NewDecimalPrecision(u)
	require.NoError(t, err)
	require.NotZero(t, v)
	return v
}

func NewDecimalTestHelper(t *testing.T, u float64) core.Decimal {
	v, err := core.NewDecimal(u)
	require.NoError(t, err)
	require.NotZero(t, v)
	return v
}
