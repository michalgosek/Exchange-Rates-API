package testabilities

import (
	"exchange-rates-api/internal/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewCalculatedExchangeRateTestHelper(t *testing.T, from, to core.CurrencyCode, val string) core.CalculatedExchangeRate {
	t.Helper()
	dec, err := core.NewExchangeRate(val)
	require.Nil(t, err)
	require.NotZero(t, dec)
	return core.NewCalculatedExchangeRate(from, to, dec)
}

func NewGlobalCurrencyExchangeRateTestHelper(t *testing.T, code core.CurrencyCode, val float64) core.GlobalCurrencyExchangeRate {
	t.Helper()
	dec, err := core.NewDecimal(val)
	require.Nil(t, err)
	require.NotZero(t, dec)
	return core.NewGlobalCurrencyExchangeRate(code, dec)
}
