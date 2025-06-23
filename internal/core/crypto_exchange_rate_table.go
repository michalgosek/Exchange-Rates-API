package core

import "fmt"

type CryptoExchangeRateTable map[CurrencyCode]CryptoExchangeRateTableEntry

func (c CryptoExchangeRateTable) AddExchangeRate(code CurrencyCode, rec CryptoExchangeRateTableEntry) {
	c[code] = rec
}

func (c CryptoExchangeRateTable) GetExchangeRate(code CurrencyCode) (CryptoExchangeRateTableEntry, error) {
	v, ok := c[code]
	if !ok {
		return CryptoExchangeRateTableEntry{}, fmt.Errorf("invalid or unsupported crypto currency code: %s", code)
	}
	return v, nil
}

type CryptoExchangeRateTableEntry struct {
	rate      ExchangeRate
	precision DecimalPrecision
}

func (c *CryptoExchangeRateTableEntry) Rate() ExchangeRate                 { return c.rate }
func (c *CryptoExchangeRateTableEntry) DecimalPrecision() DecimalPrecision { return c.precision }

func (c *CryptoExchangeRateTableEntry) EqualsTo(outer CryptoExchangeRateTableEntry) bool {
	return c.rate.EqualsTo(outer.Rate()) && c.precision.EqualsTo(outer.DecimalPrecision())
}

func NewCryptoExchangeRateTableEntry(rate ExchangeRate, prec DecimalPrecision) CryptoExchangeRateTableEntry {
	return CryptoExchangeRateTableEntry{rate: rate, precision: prec}
}

func NewDefaultCryptoExchangeRateTable() CryptoExchangeRateTable {
	return CryptoExchangeRateTable{
		BEER: CryptoExchangeRateTableEntry{
			rate:      ExchangeRate{val: "0.00002461"},
			precision: DecimalPrecision{val: 18},
		},
		FLOKI: CryptoExchangeRateTableEntry{
			rate:      ExchangeRate{val: "0.0001428"},
			precision: DecimalPrecision{val: 18},
		},
		GATE: CryptoExchangeRateTableEntry{
			rate:      ExchangeRate{val: "6.87"},
			precision: DecimalPrecision{val: 18},
		},
		// It seems that the value provided in the assignment's crypto table (i.e., 0.999)
		// was incorrect, as it does not yield WBTC → USDT = 57613.353535
		USDT: CryptoExchangeRateTableEntry{
			rate:      ExchangeRate{val: "0.990"},
			precision: DecimalPrecision{val: 6},
		},
		WBTC: CryptoExchangeRateTableEntry{
			rate:      ExchangeRate{val: "57037.22"},
			precision: DecimalPrecision{val: 8},
		},
	}
}
