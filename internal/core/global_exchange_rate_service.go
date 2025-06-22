package core

import (
	"context"
	"fmt"
)

type GlobalExchangeArithmeticService interface {
	CalculateCrossRate(first, second ExchangeRate, amount Decimal) (ExchangeRate, error)
}

type GlobalExchangeRateService struct {
	service GlobalExchangeArithmeticService
	amount  Decimal
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
			exchange1, err := g.service.CalculateCrossRate(rate2, rate1, g.amount)
			if err != nil {
				return nil, err
			}

			// from the second  rate (USD) to the first rate (EUR) i.e. USD -> EUR
			exchange2, err := g.service.CalculateCrossRate(rate1, rate2, g.amount)
			if err != nil {
				return nil, err
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

	return &GlobalExchangeRateService{
		service: service,
		amount:  Decimal{val: 1},
	}
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
