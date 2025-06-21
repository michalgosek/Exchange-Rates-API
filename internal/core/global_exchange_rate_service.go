package core

import (
	"context"
	"fmt"
)

type GlobalExchangeArithmeticService interface {
	Div(ExchangeRate, ExchangeRate) (ExchangeRate, error)
}

type GlobalExchangeRateService struct {
	service GlobalExchangeArithmeticService
}

func (g *GlobalExchangeRateService) CalculateExchangeRates(ctx context.Context, base GlobalCurrencyExchangeRate, rates ...GlobalCurrencyExchangeRate) ([]CalculatedExchangeRate, error) {
	if len(rates) < 2 {
		return nil, fmt.Errorf("invalid rates number: fewer than two rates are provided")
	}

	var exchanges []CalculatedExchangeRate

	total := append(rates, base)
	memo := make(GlobalRatePairSet)
	for _, first := range total {
		for _, second := range total {
			if first.EqualsTo(second) {
				continue
			}

			if memo.HasPair(first, second) || memo.HasPair(second, first) {
				continue
			}

			rate1, err := NewExchangeRate(first.rate.String())
			if err != nil {
				return nil, fmt.Errorf("failed to create exchange rate: %w", err)
			}

			rate2, err := NewExchangeRate(second.Rate().String())
			if err != nil {
				return nil, fmt.Errorf("failed to create exchange rate: %w", err)
			}

			// from the first rate (EUR) to the second rate (USD) i.e. EUR -> USD
			exchange1, err := g.service.Div(rate2, rate1)
			if err != nil {
				return nil, fmt.Errorf("failed to div rates (%s / %s): %w", rate2.String(), rate1.String(), err)
			}

			// from the second  rate (USD) to the first rate (EUR) i.e. USD -> EUR
			exchange2, err := g.service.Div(rate1, rate2)
			if err != nil {
				return nil, fmt.Errorf("failed to div rates (%s / %s): %w", rate1.String(), rate2.String(), err)
			}

			calculated1 := NewCalculatedExchangeRate(first.Code(), second.code, exchange1)
			calculated2 := NewCalculatedExchangeRate(second.Code(), first.Code(), exchange2)

			exchanges = append(exchanges, calculated1, calculated2)
			memo.AddPair(calculated1, calculated2)
		}
	}

	return exchanges, nil
}

func NewGlobalExchangeRateService(service GlobalExchangeArithmeticService) *GlobalExchangeRateService {
	if service == nil {
		panic("global exchange arithmetic service is required")
	}

	return &GlobalExchangeRateService{service: service}
}

type GlobalRatePairSet map[string]struct{}

func (s GlobalRatePairSet) HasPair(first, second GlobalCurrencyExchangeRate) bool {
	_, ok := s[first.Code().FromTo(second.code)]
	return ok
}

func (s GlobalRatePairSet) AddPair(rates ...CalculatedExchangeRate) {
	for _, r := range rates {
		s[r.FromTo()] = struct{}{}
	}
}
