package core

import (
	"fmt"
	"regexp"
)

var exchangeRateRegex = regexp.MustCompile(`^\d+(\.\d+)?$`)

type ExchangeRate struct{ val string }

func (e ExchangeRate) String() string                   { return e.val }
func (e ExchangeRate) EqualsTo(outer ExchangeRate) bool { return e.val == outer.val }

func NewExchangeRate(rate string) (ExchangeRate, error) {
	if !exchangeRateRegex.MatchString(rate) {
		return ExchangeRate{}, fmt.Errorf("invalid exchange rate: must be a non-negative number")
	}

	return ExchangeRate{val: rate}, nil
}

type CalculatedExchangeRate struct {
	from CurrencyCode
	to   CurrencyCode
	rate ExchangeRate
}

func (c *CalculatedExchangeRate) From() string               { return c.from.String() }
func (c *CalculatedExchangeRate) To() string                 { return c.to.String() }
func (c *CalculatedExchangeRate) FromTo() string             { return c.from.String() + "->" + c.to.String() }
func (c *CalculatedExchangeRate) ExchangeRate() ExchangeRate { return c.rate }

func NewCalculatedExchangeRate(from, to CurrencyCode, rate ExchangeRate) CalculatedExchangeRate {
	return CalculatedExchangeRate{from: from, to: to, rate: rate}
}
