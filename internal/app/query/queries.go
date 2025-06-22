package query

import "fmt"

type GlobalExchangeRatesQuery struct {
	currencies string
	base       string
}

func (g *GlobalExchangeRatesQuery) Currencies() string { return g.currencies }
func (g *GlobalExchangeRatesQuery) Base() string       { return g.base }

func NewGlobalExchangeRatesQuery(currencies, base string) (*GlobalExchangeRatesQuery, error) {
	if len(base) == 0 {
		return nil, fmt.Errorf("invalid base: must be a non empty string")
	}
	if len(currencies) == 0 {
		return nil, fmt.Errorf("invalid currencies: must be a non empty string")
	}

	return &GlobalExchangeRatesQuery{currencies: currencies, base: base}, nil
}

type CryptoExchangeRateQuery struct {
	from   string
	to     string
	amount float64
}

func (c *CryptoExchangeRateQuery) From() string    { return c.from }
func (c *CryptoExchangeRateQuery) To() string      { return c.to }
func (c *CryptoExchangeRateQuery) Amount() float64 { return c.amount }

func NewCryptoExchangeRateQuery(from, to string, amount float64) (*CryptoExchangeRateQuery, error) {
	if len(from) == 0 {
		return nil, fmt.Errorf("invalid from: must be a non empty string")
	}
	if len(to) == 0 {
		return nil, fmt.Errorf("invalid to: must be a non empty string")
	}
	if amount < 0 {
		return nil, fmt.Errorf("invalid amount: must be a non negative number greater than zero")
	}

	return &CryptoExchangeRateQuery{from: from, to: to, amount: amount}, nil
}
