package math_test

import (
	"exchange-rates-api/internal/core"
	"exchange-rates-api/internal/infrastructure/math"
	"exchange-rates-api/internal/infrastructure/math/testabilities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrencyExchangeArithmeticService_CalculateCrossRateWithPrecission(t *testing.T) {
	tests := map[string]struct {
		expectedRate string
		from         core.ExchangeRate
		to           core.ExchangeRate
		prec         core.DecimalPrecision
		amout        core.Decimal
	}{
		"should provide exchange rate from USDT with 8 dec places to WBTC with 6 dec places": {
			expectedRate: "57613.35353535",
			from:         testabilities.NewExchangeRateTestHelper(t, "57037.22"),
			to:           testabilities.NewExchangeRateTestHelper(t, "0.990"),
			prec:         testabilities.NewDecimalPrecisionTestHelper(t, 8),
			amout:        testabilities.NewDecimalTestHelper(t, 1),
		},
		"should provide exchange rate from WBTC with 6 dec places to USDT with 8 dec places": {
			expectedRate: "0.000017",
			from:         testabilities.NewExchangeRateTestHelper(t, "0.990"),
			to:           testabilities.NewExchangeRateTestHelper(t, "57037.22"),
			prec:         testabilities.NewDecimalPrecisionTestHelper(t, 6),
			amout:        testabilities.NewDecimalTestHelper(t, 1),
		},
		"should provide exchange rate from GATE with 18 dec places to FLOKI with 18 dec places": {
			expectedRate: "0.000020786026200873",
			from:         testabilities.NewExchangeRateTestHelper(t, "0.0001428"),
			to:           testabilities.NewExchangeRateTestHelper(t, "6.87"),
			prec:         testabilities.NewDecimalPrecisionTestHelper(t, 18),
			amout:        testabilities.NewDecimalTestHelper(t, 1),
		},
		"should provide exchange rate from BEER with 18 dec places to WBTC with 8 dec places": {
			expectedRate: "0.00000000",
			from:         testabilities.NewExchangeRateTestHelper(t, "0.00002461"),
			to:           testabilities.NewExchangeRateTestHelper(t, "57037.22"),
			prec:         testabilities.NewDecimalPrecisionTestHelper(t, 8),
			amout:        testabilities.NewDecimalTestHelper(t, 1),
		},
		"should provide exchange rate from WBTC with 8 dec places to BEER with 18 dec places": {
			expectedRate: "2317644047.135310849248273060",
			from:         testabilities.NewExchangeRateTestHelper(t, "57037.22"),
			to:           testabilities.NewExchangeRateTestHelper(t, "0.00002461"),
			prec:         testabilities.NewDecimalPrecisionTestHelper(t, 18),
			amout:        testabilities.NewDecimalTestHelper(t, 1),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// given:
			service := math.NewCurrencyExchangeArithmeticService(math.DefaultExchangeRatePrecision)

			// when:
			actualRate, err := service.CalculateCrossRateWithPrecission(tc.from, tc.to, tc.amout, tc.prec)

			// then:
			require.Nil(t, err)
			require.Equal(t, tc.expectedRate, actualRate.String())
		})
	}
}
