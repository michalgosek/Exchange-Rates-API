package core

import "fmt"

var (
	BEER  = CurrencyCode{"BEER"}
	FLOKI = CurrencyCode{"FLOKI"}
	GATE  = CurrencyCode{"GATE"}
	USDT  = CurrencyCode{"USDT"}
	WBTC  = CurrencyCode{"WBTC"}
)

var cryptoCurrencyCodesRegister = map[string]CurrencyCode{
	"BEER":  BEER,
	"FLOKI": FLOKI,
	"GATE":  GATE,
	"USDT":  USDT,
	"WBTC":  WBTC,
}

func NewCryptoCurrencyCode(code string) (CurrencyCode, error) {
	val, ok := cryptoCurrencyCodesRegister[code]
	if !ok {
		return CurrencyCode{}, fmt.Errorf("invalid or unsupported crypto currency code: %s", code)
	}
	return val, nil
}
