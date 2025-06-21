package adapters

import (
	"exchange-rates-api/internal/core"
	"fmt"
)

type ExchangeRatesDTO struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func (e ExchangeRatesDTO) ToGlobalCurrencyExchangeRate() ([]core.GlobalCurrencyExchangeRate, error) {
	var rates []core.GlobalCurrencyExchangeRate
	for k, v := range e.Rates {
		code, err := core.NewGlobalCurrencyCode(k)
		if err != nil {
			return nil, fmt.Errorf("failed to create global currency code: %w", err)
		}
		dec, err := core.NewDecimal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create decimal: %w", err)
		}
		rates = append(rates, core.NewGlobalCurrencyExchangeRate(code, dec))
	}

	return rates, nil
}
