package app

import "exchange-rates-api/internal/app/query"

type Queries struct {
	GlobalExchangeRatesHandler *query.GlobalExchangeRatesHandler
	CryptoExchangeRateHandler  *query.CryptoExchangeRateHandler
}

type Application struct {
	Queries Queries
}
