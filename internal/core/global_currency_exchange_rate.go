package core

type GlobalCurrencyExchangeRate struct {
	code CurrencyCode
	rate Decimal
}

func (g GlobalCurrencyExchangeRate) Rate() Decimal      { return g.rate }
func (g GlobalCurrencyExchangeRate) Code() CurrencyCode { return g.code }

func (g GlobalCurrencyExchangeRate) EqualsTo(outer GlobalCurrencyExchangeRate) bool {
	return g.Code().EqualsTo(outer.Code()) && g.rate.EqualsTo(outer.Rate())
}

func NewGlobalCurrencyExchangeRate(code CurrencyCode, rate Decimal) GlobalCurrencyExchangeRate {
	return GlobalCurrencyExchangeRate{code: code, rate: rate}
}
