package adapters

import (
	"context"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type OpenExchangeRatesHTTPConfig struct {
	AppID   string `mapstructure:"app_id"`
	BaseURL string `mapstructure:"base_url"`
}

type OpenExchangeRatesHTTP struct {
	cfg  *OpenExchangeRatesHTTPConfig
	http *resty.Client
}

func (o *OpenExchangeRatesHTTP) GetLatestExchangeRates(ctx context.Context, query query.GlobalExchangeRatesQuery) ([]core.GlobalCurrencyExchangeRate, error) {
	const URL = "/latest.json"

	var dto ExchangeRatesDTO
	res, err := o.http.
		R().
		SetContext(ctx).
		SetResult(&dto).
		SetQueryParams(map[string]string{
			"base":    query.Base(),
			"symbols": query.Currencies(),
		}).
		Get("/latest.json")

	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP GET req to url %s: %w", o.cfg.BaseURL+URL, err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("failed to send HTTP GET req to url %s. Status code: %s", o.cfg.BaseURL+URL, res.Status())
	}

	return dto.ToGlobalCurrencyExchangeRate()
}

// TODO: Recommended optimization to avoid wasting available request quota depending on subscription -> https://docs.openexchangerates.org/reference/etags
// TODO: Here we can add a more robust implementation in case of API failure :)
func NewOpenExchangeRatesHTTP(cfg *OpenExchangeRatesHTTPConfig) *OpenExchangeRatesHTTP {
	http := resty.
		New().
		SetBaseURL(cfg.BaseURL).
		SetRetryCount(5).
		SetHeader("Authorization", "Token "+cfg.AppID)

	return &OpenExchangeRatesHTTP{cfg: cfg, http: http}
}
