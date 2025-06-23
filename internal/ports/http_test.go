package ports_test

import (
	"exchange-rates-api/internal/app"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"exchange-rates-api/internal/infrastructure/math"
	"exchange-rates-api/internal/ports"
	ports_testabilities "exchange-rates-api/internal/ports/testabilities"

	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPServer_GetGlobalExchangeRates_NegativeTestCases(t *testing.T) {
	tests := map[string]struct {
		params                            url.Values
		expectedTestGlobalRatesServiceCfg ports_testabilities.TestGlobalRatesServiceConfig
	}{
		"should respond with 400 status code when required query parameters are missing in /api/v1/rates request - an empty query params list": {
			params:                            url.Values{},
			expectedTestGlobalRatesServiceCfg: ports_testabilities.TestGlobalRatesServiceConfig{ExpectedGetLatestExchangeRatesCall: false},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/rates request - base query param only": {
			params:                            url.Values{"base": []string{"USD"}},
			expectedTestGlobalRatesServiceCfg: ports_testabilities.TestGlobalRatesServiceConfig{ExpectedGetLatestExchangeRatesCall: false},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/rates request - query params list with base and one currency": {
			params: url.Values{
				"base":       []string{"USD"},
				"currencies": []string{"PLN"},
			},
			expectedTestGlobalRatesServiceCfg: ports_testabilities.TestGlobalRatesServiceConfig{ExpectedGetLatestExchangeRatesCall: true},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// given:
			mock := ports_testabilities.NewTestGlobalRatesService(t, tc.expectedTestGlobalRatesServiceCfg)
			fixture := ports.NewHTTPServerTestFixture(t, &app.Application{
				Queries: &app.Queries{
					GlobalExchangeRatesHandler: query.NewGlobalExchangeRatesHandler(
						mock,
						core.NewGlobalExchangeRateService(math.NewCurrencyExchangeArithmeticService(math.GlobalCurrencyExchangeRatePrecision)),
					),
				},
			})

			// when:
			rec := fixture.NewResponseRecorder()
			req := fixture.NewRequestWithParams(t.Context(), http.MethodGet, "/api/v1/rates", tc.params)
			fixture.ServeHTTP(req, rec)

			// then:
			resp := rec.Result()
			expectedStatusCode := http.StatusBadRequest

			require.Equal(t, expectedStatusCode, resp.StatusCode)
			mock.AssertCalled()
		})
	}
}

func TestHTTPServer_GetGlobalExchangeRates_PostivieTestCase(t *testing.T) {
	t.Run("should respond with calculated exchange rates rounded to 5 decimal places for /api/v1/rates when queried with base=USD and currencies=EUR,PLN", func(t *testing.T) {
		// given:
		mock := ports_testabilities.NewTestGlobalRatesService(t, ports_testabilities.DefaultGlobalTestRatesServiceConfig)
		fixture := ports.NewHTTPServerTestFixture(t, &app.Application{
			Queries: &app.Queries{
				GlobalExchangeRatesHandler: query.NewGlobalExchangeRatesHandler(
					mock,
					core.NewGlobalExchangeRateService(math.NewCurrencyExchangeArithmeticService(math.GlobalCurrencyExchangeRatePrecision)),
				),
			},
		})

		params := make(url.Values)
		params.Add("currencies", "EUR,PLN")
		params.Add("base", "USD")

		expectedStatusCode := http.StatusOK
		expectedDTOs := []ports.GlobalExchangeRateDTO{
			{
				From: "EUR",
				To:   "PLN",
				Rate: "4.2686",
			},
			{
				From: "PLN",
				To:   "EUR",
				Rate: "0.23427",
			},
		}

		// when:
		rec := fixture.NewResponseRecorder()
		req := fixture.NewRequestWithParams(t.Context(), http.MethodGet, "/api/v1/rates", params)
		fixture.ServeHTTP(req, rec)

		// then:
		var actualDTOs []ports.GlobalExchangeRateDTO

		resp := rec.Result()
		fixture.DecodeResponse(resp, &actualDTOs)

		require.Equal(t, expectedDTOs, actualDTOs)
		require.Equal(t, expectedStatusCode, resp.StatusCode)
		mock.AssertCalled()
	})
}

func TestHTTPServer_GetCryptoExchangeRate_NegativeTestCases(t *testing.T) {
	tests := map[string]struct {
		params url.Values
	}{
		"should respond with 400 status code when required query parameters are missing in /api/v1/convert request - an empty query params list": {
			params: url.Values{},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/convert request - from query param only": {
			params: url.Values{"from": []string{"WBTC"}},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/convert request - to query param only": {
			params: url.Values{"to": []string{"WBTC"}},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/convert request - amount query param only": {
			params: url.Values{"amount": []string{"1"}},
		},
		"should respond with 400 status code when required query parameters are missing in /api/v1/convert request - query params list with from and to currency": {
			params: url.Values{
				"from": []string{"USDT"},
				"to":   []string{"WBTC"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// given:
			service := core.NewCryptoExchangeRateService(math.NewCurrencyExchangeArithmeticService(math.GlobalCurrencyExchangeRatePrecision), core.CryptoExchangeRateTable{})
			fixture := ports.NewHTTPServerTestFixture(t, &app.Application{
				Queries: &app.Queries{
					CryptoExchangeRateHandler: query.NewCryptoExchangeRateHandler(service),
				},
			})

			rec := fixture.NewResponseRecorder()
			req := fixture.NewRequestWithParams(t.Context(), http.MethodGet, "/api/v1/exchange", tc.params)
			expectedStatusCode := http.StatusBadRequest

			// when:
			fixture.ServeHTTP(req, rec)

			// then:
			resp := rec.Result()
			require.Equal(t, expectedStatusCode, resp.StatusCode)
		})
	}
}

func TestHTTPServer_GetCryptoExchangeRate_PositiveTestCase(t *testing.T) {
	t.Run("should respond with calculated exchange rate rounded to 5 decimal places for /api/v1/convert when queried with from=WBTC, to=USDT and amount=1", func(t *testing.T) {
		// given:
		const WBTCPrecision = 5
		const USDTPrecision = 6

		table := make(core.CryptoExchangeRateTable)
		table.AddExchangeRate(core.USDT, core.NewCryptoExchangeRateTableEntry(ports_testabilities.NewExchangeRateTestHelper(t, "0.990"), ports_testabilities.NewDecimalPrecisionTestHelper(t, USDTPrecision)))
		table.AddExchangeRate(core.WBTC, core.NewCryptoExchangeRateTableEntry(ports_testabilities.NewExchangeRateTestHelper(t, "57037.22"), ports_testabilities.NewDecimalPrecisionTestHelper(t, WBTCPrecision)))
		service := core.NewCryptoExchangeRateService(math.NewCurrencyExchangeArithmeticService(math.DefaultExchangeRatePrecision), table)

		fixture := ports.NewHTTPServerTestFixture(t, &app.Application{
			Queries: &app.Queries{
				CryptoExchangeRateHandler: query.NewCryptoExchangeRateHandler(service),
			},
		})

		params := make(url.Values)
		params.Add("from", "WBTC")
		params.Add("to", "USDT")
		params.Add("amount", "1")

		rec := fixture.NewResponseRecorder()
		req := fixture.NewRequestWithParams(t.Context(), http.MethodGet, "/api/v1/exchange", params)
		expectedStatusCode := http.StatusOK
		expectedDTO := ports.CryptoExchangeRateDTO{
			From:   "WBTC",
			Amount: "57613.353535",
			To:     "USDT",
		}

		// when:
		fixture.ServeHTTP(req, rec)

		// then:
		var actualDTO ports.CryptoExchangeRateDTO

		resp := rec.Result()
		fixture.DecodeResponse(resp, &actualDTO)

		require.Equal(t, expectedDTO, actualDTO)
		require.Equal(t, expectedStatusCode, resp.StatusCode)
	})
}
