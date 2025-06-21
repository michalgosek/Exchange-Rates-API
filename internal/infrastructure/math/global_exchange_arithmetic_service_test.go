package math_test

import (
	"exchange-rates-api/internal/infrastructure/math"
	"exchange-rates-api/internal/infrastructure/math/testabilities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobalExchangeArithmeticService_DivWithPrecision(t *testing.T) {
	t.Run("should provide exchange rate from USDT with 8 dec places to WBTC with 6 dec places", func(t *testing.T) {
		// given:
		const ExpectedUSDTRate = "57613.35353535"
		const WBTCPrecision = 8

		wbtc := testabilities.NewExchangeRateTestHelper(t, "57037.22")
		usdt := testabilities.NewExchangeRateTestHelper(t, "0.990")
		prec := testabilities.NewDecimalPrecisionTestHelper(t, WBTCPrecision)
		service := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)

		// when:
		actualRate, err := service.DivWithPrecision(wbtc, usdt, prec)

		// then:
		require.Nil(t, err)
		require.Equal(t, ExpectedUSDTRate, actualRate.String())
	})

	t.Run("should provide exchange rate from WBTC with 6 dec places to USDT with 8 dec places", func(t *testing.T) {
		// given:
		const ExpectedWBTCRate = "0.000017"
		const USDTPrecision = 6

		wbtc := testabilities.NewExchangeRateTestHelper(t, "57037.22")
		usdt := testabilities.NewExchangeRateTestHelper(t, "0.990")
		prec := testabilities.NewDecimalPrecisionTestHelper(t, USDTPrecision)
		service := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)

		// when:
		actualRate, err := service.DivWithPrecision(usdt, wbtc, prec)

		// then:
		require.Nil(t, err)
		require.Equal(t, ExpectedWBTCRate, actualRate.String())
	})

	t.Run("should provide exchange rate from GATE with 18 dec places to FLOKI with 18 dec places", func(t *testing.T) {
		// given:
		const ExpectedFLOKIRate = "0.000020786026200873"
		const Precision = 18

		floki := testabilities.NewExchangeRateTestHelper(t, "0.0001428")
		gate := testabilities.NewExchangeRateTestHelper(t, "6.87")
		prec := testabilities.NewDecimalPrecisionTestHelper(t, Precision)
		service := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)

		// when:
		actualRate, err := service.DivWithPrecision(floki, gate, prec)

		// then:
		require.Nil(t, err)
		require.Equal(t, ExpectedFLOKIRate, actualRate.String())
	})

	t.Run("should provide exchange rate from BEER with 18 dec places to WBTC with 8 dec places", func(t *testing.T) {
		// given:
		const ExpectedWBTCRate = "0.00000000"
		const WBTCPrecision = 8

		beer := testabilities.NewExchangeRateTestHelper(t, "0.00002461")
		wbtc := testabilities.NewExchangeRateTestHelper(t, "57037.22")
		prec := testabilities.NewDecimalPrecisionTestHelper(t, WBTCPrecision)
		service := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)

		// when:
		actualRate, err := service.DivWithPrecision(beer, wbtc, prec)

		// then:
		require.Nil(t, err)
		require.Equal(t, ExpectedWBTCRate, actualRate.String())
	})

	t.Run("should provide exchange rate from WBTC with 8 dec places to BEER with 18 dec places", func(t *testing.T) {
		// given:
		const ExpectedBEERRate = "2317644047.135310849248273060"
		const BEERPrecision = 18

		beer := testabilities.NewExchangeRateTestHelper(t, "0.00002461")
		wbtc := testabilities.NewExchangeRateTestHelper(t, "57037.22")
		prec := testabilities.NewDecimalPrecisionTestHelper(t, BEERPrecision)
		service := math.NewGlobalExchangeArithmeticService(math.IEE754Precision)

		// when:
		actualRate, err := service.DivWithPrecision(wbtc, beer, prec)

		// then:
		require.Nil(t, err)
		require.Equal(t, ExpectedBEERRate, actualRate.String())
	})
}
