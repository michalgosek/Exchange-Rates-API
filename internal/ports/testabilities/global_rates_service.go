package testabilities

import (
	"context"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"testing"

	"github.com/stretchr/testify/require"
)

var DefaultGlobalTestRatesServiceConfig = TestGlobalRatesServiceConfig{
	ExpectedErr:                        nil,
	ExpectedGetLatestExchangeRatesCall: true,
	Rates: map[string]float64{
		"EUR": 0.867717,
		"PLN": 3.703935,
	},
}

type TestGlobalRatesServiceConfig struct {
	ExpectedErr                        error
	ExpectedGetLatestExchangeRatesCall bool
	Rates                              map[string]float64
}

type TestGlobalRatesService struct {
	t      *testing.T
	cfg    TestGlobalRatesServiceConfig
	called bool
}

func (g *TestGlobalRatesService) AssertCalled() {
	g.t.Helper()
	require.Equal(g.t, g.cfg.ExpectedGetLatestExchangeRatesCall, g.called)
}

func (g *TestGlobalRatesService) GetLatestExchangeRates(ctx context.Context, query query.GlobalExchangeRatesQuery) ([]core.GlobalCurrencyExchangeRate, error) {
	g.called = true

	g.t.Helper()
	if g.cfg.ExpectedErr != nil {
		return nil, g.cfg.ExpectedErr
	}

	var rates []core.GlobalCurrencyExchangeRate
	for k, v := range g.cfg.Rates {
		code, err := core.NewGlobalCurrencyCode(k)
		require.NoError(g.t, err)
		require.NotEmpty(g.t, code)

		dec, err := core.NewDecimal(v)
		require.NoError(g.t, err)
		require.NotEmpty(g.t, dec)
		rates = append(rates, core.NewGlobalCurrencyExchangeRate(code, dec))
	}

	return rates, nil
}

func NewTestGlobalRatesService(t *testing.T, cfg TestGlobalRatesServiceConfig) *TestGlobalRatesService {
	return &TestGlobalRatesService{t: t, cfg: cfg}
}
