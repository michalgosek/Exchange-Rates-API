package core_test

import (
	"exchange-rates-api/internal/core"
	"exchange-rates-api/internal/core/testabilities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobalExchangeRateService_CalculateExchangeRates(t *testing.T) {
	t.Run("should return an error when fewer than two rates are provided", func(t *testing.T) {
		// given:
		const DecimalPrecision = 5

		usd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.USD, 1)
		pln := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.PLN, 3.703935)

		arithmetic := testabilities.NewTestGlobalExchangeArithmeticServiceFloat64(t, DecimalPrecision)
		service := core.NewGlobalExchangeRateService(arithmetic)

		// when:
		actualRates, err := service.CalculateExchangeRates(t.Context(), usd, pln)

		// then:
		require.NotNil(t, err) // We can also apply more robust approach for handling error :)
		require.Nil(t, actualRates)
	})

	t.Run("should provide exchange rates for both EUR, PLN, based on USD", func(t *testing.T) {
		// given:
		const DecimalPrecision = 5

		usd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.USD, 1)
		eur := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.EUR, 0.868973)
		pln := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.PLN, 3.703935)

		arithmetic := testabilities.NewTestGlobalExchangeArithmeticServiceFloat64(t, DecimalPrecision)
		service := core.NewGlobalExchangeRateService(arithmetic)

		expectedRates := []core.CalculatedExchangeRate{
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.EUR, core.PLN, "4.26243"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.PLN, core.EUR, "0.23461"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.EUR, core.USD, "1.15078"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.EUR, "0.86897"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.PLN, core.USD, "0.26998"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.PLN, "3.70394"),
		}

		// when:
		actualRates, err := service.CalculateExchangeRates(t.Context(), usd, eur, pln)

		// then:
		require.NoError(t, err)
		require.Equal(t, expectedRates, actualRates)
	})

	t.Run("should provide exchange rates for both EUR, GBP based on USD", func(t *testing.T) {
		// given:
		const DecimalPrecision = 5

		usd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.USD, 1)
		eur := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.EUR, 0.868973)
		gpb := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.GBP, 0.743437)

		arithmetic := testabilities.NewTestGlobalExchangeArithmeticServiceFloat64(t, DecimalPrecision)
		service := core.NewGlobalExchangeRateService(arithmetic)

		expectedRates := []core.CalculatedExchangeRate{
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.EUR, core.GBP, "0.85554"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.GBP, core.EUR, "1.16886"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.EUR, core.USD, "1.15078"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.EUR, "0.86897"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.GBP, core.USD, "1.3451"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.GBP, "0.74344"),
		}

		// when:
		actualRates, err := service.CalculateExchangeRates(t.Context(), usd, eur, gpb)

		// then:
		require.NoError(t, err)
		require.Equal(t, expectedRates, actualRates)
	})

	t.Run("should provide exchange rates for both ARS, AOA based on USD", func(t *testing.T) {
		// given:
		const DecimalPrecision = 5

		usd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.USD, 1)
		aoa := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.AOA, 911.955)
		ars := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.ARS, 1164.51527)

		arithmetic := testabilities.NewTestGlobalExchangeArithmeticServiceFloat64(t, DecimalPrecision)
		service := core.NewGlobalExchangeRateService(arithmetic)

		expectedRates := []core.CalculatedExchangeRate{
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.AOA, core.ARS, "1.27694"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.ARS, core.AOA, "0.78312"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.AOA, core.USD, "0.0011"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.AOA, "911.955"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.ARS, core.USD, "0.00086"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.ARS, "1164.51527"),
		}

		// when:
		actualRates, err := service.CalculateExchangeRates(t.Context(), usd, aoa, ars)

		// then:
		require.NoError(t, err)
		require.Equal(t, expectedRates, actualRates)
	})

	t.Run("should provide exchange rates for both DOP, DZD based on USD", func(t *testing.T) {
		// given:
		const DecimalPrecision = 5

		usd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.USD, 1)
		dzd := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.DZD, 130.222172)
		dop := testabilities.NewGlobalCurrencyExchangeRateTestHelper(t, core.DOP, 59.020702)

		arithmetic := testabilities.NewTestGlobalExchangeArithmeticServiceFloat64(t, DecimalPrecision)
		service := core.NewGlobalExchangeRateService(arithmetic)

		expectedRates := []core.CalculatedExchangeRate{
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.DZD, core.DOP, "0.45323"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.DOP, core.DZD, "2.20638"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.DZD, core.USD, "0.00768"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.DZD, "130.22217"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.DOP, core.USD, "0.01694"),
			testabilities.NewCalculatedExchangeRateTestHelper(t, core.USD, core.DOP, "59.0207"),
		}

		// when:
		actualRates, err := service.CalculateExchangeRates(t.Context(), usd, dzd, dop)

		// then:
		require.NoError(t, err)
		require.Equal(t, expectedRates, actualRates)
	})
}
