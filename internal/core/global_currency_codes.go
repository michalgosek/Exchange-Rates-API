package core

import "fmt"

var (
	AED = CurrencyCode{"AED"}
	AFN = CurrencyCode{"AFN"}
	ALL = CurrencyCode{"ALL"}
	AMD = CurrencyCode{"AMD"}
	ANG = CurrencyCode{"ANG"}
	AOA = CurrencyCode{"AOA"}
	ARS = CurrencyCode{"ARS"}
	AUD = CurrencyCode{"AUD"}
	AWG = CurrencyCode{"AWG"}
	AZN = CurrencyCode{"AZN"}
	BAM = CurrencyCode{"BAM"}
	BBD = CurrencyCode{"BBD"}
	BDT = CurrencyCode{"BDT"}
	BGN = CurrencyCode{"BGN"}
	BHD = CurrencyCode{"BHD"}
	BIF = CurrencyCode{"BIF"}
	BMD = CurrencyCode{"BMD"}
	BND = CurrencyCode{"BND"}
	BOB = CurrencyCode{"BOB"}
	BRL = CurrencyCode{"BRL"}
	BSD = CurrencyCode{"BSD"}
	BTC = CurrencyCode{"BTC"}
	BTN = CurrencyCode{"BTN"}
	BWP = CurrencyCode{"BWP"}
	BYN = CurrencyCode{"BYN"}
	BZD = CurrencyCode{"BZD"}
	CAD = CurrencyCode{"CAD"}
	CDF = CurrencyCode{"CDF"}
	CHF = CurrencyCode{"CHF"}
	CLF = CurrencyCode{"CLF"}
	CLP = CurrencyCode{"CLP"}
	CNH = CurrencyCode{"CNH"}
	CNY = CurrencyCode{"CNY"}
	COP = CurrencyCode{"COP"}
	CRC = CurrencyCode{"CRC"}
	CUC = CurrencyCode{"CUC"}
	CUP = CurrencyCode{"CUP"}
	CVE = CurrencyCode{"CVE"}
	CZK = CurrencyCode{"CZK"}
	DJF = CurrencyCode{"DJF"}
	DKK = CurrencyCode{"DKK"}
	DOP = CurrencyCode{"DOP"}
	DZD = CurrencyCode{"DZD"}
	EGP = CurrencyCode{"EGP"}
	ERN = CurrencyCode{"ERN"}
	ETB = CurrencyCode{"ETB"}
	EUR = CurrencyCode{"EUR"}
	FJD = CurrencyCode{"FJD"}
	FKP = CurrencyCode{"FKP"}
	GBP = CurrencyCode{"GBP"}
	GEL = CurrencyCode{"GEL"}
	GGP = CurrencyCode{"GGP"}
	GHS = CurrencyCode{"GHS"}
	GIP = CurrencyCode{"GIP"}
	GMD = CurrencyCode{"GMD"}
	GNF = CurrencyCode{"GNF"}
	GTQ = CurrencyCode{"GTQ"}
	GYD = CurrencyCode{"GYD"}
	HKD = CurrencyCode{"HKD"}
	HNL = CurrencyCode{"HNL"}
	HRK = CurrencyCode{"HRK"}
	HTG = CurrencyCode{"HTG"}
	HUF = CurrencyCode{"HUF"}
	IDR = CurrencyCode{"IDR"}
	ILS = CurrencyCode{"ILS"}
	IMP = CurrencyCode{"IMP"}
	INR = CurrencyCode{"INR"}
	IQD = CurrencyCode{"IQD"}
	IRR = CurrencyCode{"IRR"}
	ISK = CurrencyCode{"ISK"}
	JEP = CurrencyCode{"JEP"}
	JMD = CurrencyCode{"JMD"}
	JOD = CurrencyCode{"JOD"}
	JPY = CurrencyCode{"JPY"}
	KES = CurrencyCode{"KES"}
	KGS = CurrencyCode{"KGS"}
	KHR = CurrencyCode{"KHR"}
	KMF = CurrencyCode{"KMF"}
	KPW = CurrencyCode{"KPW"}
	KRW = CurrencyCode{"KRW"}
	KWD = CurrencyCode{"KWD"}
	KYD = CurrencyCode{"KYD"}
	KZT = CurrencyCode{"KZT"}
	LAK = CurrencyCode{"LAK"}
	LBP = CurrencyCode{"LBP"}
	LKR = CurrencyCode{"LKR"}
	LRD = CurrencyCode{"LRD"}
	LSL = CurrencyCode{"LSL"}
	LYD = CurrencyCode{"LYD"}
	MAD = CurrencyCode{"MAD"}
	MDL = CurrencyCode{"MDL"}
	MGA = CurrencyCode{"MGA"}
	MKD = CurrencyCode{"MKD"}
	MMK = CurrencyCode{"MMK"}
	MNT = CurrencyCode{"MNT"}
	MOP = CurrencyCode{"MOP"}
	MRU = CurrencyCode{"MRU"}
	MUR = CurrencyCode{"MUR"}
	MVR = CurrencyCode{"MVR"}
	MWK = CurrencyCode{"MWK"}
	MXN = CurrencyCode{"MXN"}
	MYR = CurrencyCode{"MYR"}
	MZN = CurrencyCode{"MZN"}
	NAD = CurrencyCode{"NAD"}
	NGN = CurrencyCode{"NGN"}
	NIO = CurrencyCode{"NIO"}
	NOK = CurrencyCode{"NOK"}
	NPR = CurrencyCode{"NPR"}
	NZD = CurrencyCode{"NZD"}
	OMR = CurrencyCode{"OMR"}
	PAB = CurrencyCode{"PAB"}
	PEN = CurrencyCode{"PEN"}
	PGK = CurrencyCode{"PGK"}
	PHP = CurrencyCode{"PHP"}
	PKR = CurrencyCode{"PKR"}
	PLN = CurrencyCode{"PLN"}
	PYG = CurrencyCode{"PYG"}
	QAR = CurrencyCode{"QAR"}
	RON = CurrencyCode{"RON"}
	RSD = CurrencyCode{"RSD"}
	RUB = CurrencyCode{"RUB"}
	RWF = CurrencyCode{"RWF"}
	SAR = CurrencyCode{"SAR"}
	SBD = CurrencyCode{"SBD"}
	SCR = CurrencyCode{"SCR"}
	SDG = CurrencyCode{"SDG"}
	SEK = CurrencyCode{"SEK"}
	SGD = CurrencyCode{"SGD"}
	SHP = CurrencyCode{"SHP"}
	SLL = CurrencyCode{"SLL"}
	SOS = CurrencyCode{"SOS"}
	SRD = CurrencyCode{"SRD"}
	SSP = CurrencyCode{"SSP"}
	STD = CurrencyCode{"STD"}
	STN = CurrencyCode{"STN"}
	SVC = CurrencyCode{"SVC"}
	SYP = CurrencyCode{"SYP"}
	SZL = CurrencyCode{"SZL"}
	THB = CurrencyCode{"THB"}
	TJS = CurrencyCode{"TJS"}
	TMT = CurrencyCode{"TMT"}
	TND = CurrencyCode{"TND"}
	TOP = CurrencyCode{"TOP"}
	TRY = CurrencyCode{"TRY"}
	TTD = CurrencyCode{"TTD"}
	TWD = CurrencyCode{"TWD"}
	TZS = CurrencyCode{"TZS"}
	UAH = CurrencyCode{"UAH"}
	UGX = CurrencyCode{"UGX"}
	USD = CurrencyCode{"USD"}
	UYU = CurrencyCode{"UYU"}
	UZS = CurrencyCode{"UZS"}
	VES = CurrencyCode{"VES"}
	VND = CurrencyCode{"VND"}
	VUV = CurrencyCode{"VUV"}
	WST = CurrencyCode{"WST"}
	XAF = CurrencyCode{"XAF"}
	XAG = CurrencyCode{"XAG"}
	XAU = CurrencyCode{"XAU"}
	XCD = CurrencyCode{"XCD"}
	XDR = CurrencyCode{"XDR"}
	XOF = CurrencyCode{"XOF"}
	XPD = CurrencyCode{"XPD"}
	XPF = CurrencyCode{"XPF"}
	XPT = CurrencyCode{"XPT"}
	YER = CurrencyCode{"YER"}
	ZAR = CurrencyCode{"ZAR"}
	ZMW = CurrencyCode{"ZMW"}
	ZWL = CurrencyCode{"ZWL"}
)

var currencyCodesRegister = map[string]CurrencyCode{
	"AED": AED,
	"AFN": AFN,
	"ALL": ALL,
	"AMD": AMD,
	"ANG": ANG,
	"AOA": AOA,
	"ARS": ARS,
	"AUD": AUD,
	"AWG": AWG,
	"AZN": AZN,
	"BAM": BAM,
	"BBD": BBD,
	"BDT": BDT,
	"BGN": BGN,
	"BHD": BHD,
	"BIF": BIF,
	"BMD": BMD,
	"BND": BND,
	"BOB": BOB,
	"BRL": BRL,
	"BSD": BSD,
	"BTC": BTC,
	"BTN": BTN,
	"BWP": BWP,
	"BYN": BYN,
	"BZD": BZD,
	"CAD": CAD,
	"CDF": CDF,
	"CHF": CHF,
	"CLF": CLF,
	"CLP": CLP,
	"CNH": CNH,
	"CNY": CNY,
	"COP": COP,
	"CRC": CRC,
	"CUC": CUC,
	"CUP": CUP,
	"CVE": CVE,
	"CZK": CZK,
	"DJF": DJF,
	"DKK": DKK,
	"DOP": DOP,
	"DZD": DZD,
	"EGP": EGP,
	"ERN": ERN,
	"ETB": ETB,
	"EUR": EUR,
	"FJD": FJD,
	"FKP": FKP,
	"GBP": GBP,
	"GEL": GEL,
	"GGP": GGP,
	"GHS": GHS,
	"GIP": GIP,
	"GMD": GMD,
	"GNF": GNF,
	"GTQ": GTQ,
	"GYD": GYD,
	"HKD": HKD,
	"HNL": HNL,
	"HRK": HRK,
	"HTG": HTG,
	"HUF": HUF,
	"IDR": IDR,
	"ILS": ILS,
	"IMP": IMP,
	"INR": INR,
	"IQD": IQD,
	"IRR": IRR,
	"ISK": ISK,
	"JEP": JEP,
	"JMD": JMD,
	"JOD": JOD,
	"JPY": JPY,
	"KES": KES,
	"KGS": KGS,
	"KHR": KHR,
	"KMF": KMF,
	"KPW": KPW,
	"KRW": KRW,
	"KWD": KWD,
	"KYD": KYD,
	"KZT": KZT,
	"LAK": LAK,
	"LBP": LBP,
	"LKR": LKR,
	"LRD": LRD,
	"LSL": LSL,
	"LYD": LYD,
	"MAD": MAD,
	"MDL": MDL,
	"MGA": MGA,
	"MKD": MKD,
	"MMK": MMK,
	"MNT": MNT,
	"MOP": MOP,
	"MRU": MRU,
	"MUR": MUR,
	"MVR": MVR,
	"MWK": MWK,
	"MXN": MXN,
	"MYR": MYR,
	"MZN": MZN,
	"NAD": NAD,
	"NGN": NGN,
	"NIO": NIO,
	"NOK": NOK,
	"NPR": NPR,
	"NZD": NZD,
	"OMR": OMR,
	"PAB": PAB,
	"PEN": PEN,
	"PGK": PGK,
	"PHP": PHP,
	"PKR": PKR,
	"PLN": PLN,
	"PYG": PYG,
	"QAR": QAR,
	"RON": RON,
	"RSD": RSD,
	"RUB": RUB,
	"RWF": RWF,
	"SAR": SAR,
	"SBD": SBD,
	"SCR": SCR,
	"SDG": SDG,
	"SEK": SEK,
	"SGD": SGD,
	"SHP": SHP,
	"SLL": SLL,
	"SOS": SOS,
	"SRD": SRD,
	"SSP": SSP,
	"STD": STD,
	"STN": STN,
	"SVC": SVC,
	"SYP": SYP,
	"SZL": SZL,
	"THB": THB,
	"TJS": TJS,
	"TMT": TMT,
	"TND": TND,
	"TOP": TOP,
	"TRY": TRY,
	"TTD": TTD,
	"TWD": TWD,
	"TZS": TZS,
	"UAH": UAH,
	"UGX": UGX,
	"USD": USD,
	"UYU": UYU,
	"UZS": UZS,
	"VES": VES,
	"VND": VND,
	"VUV": VUV,
	"WST": WST,
	"XAF": XAF,
	"XAG": XAG,
	"XAU": XAU,
	"XCD": XCD,
	"XDR": XDR,
	"XOF": XOF,
	"XPD": XPD,
	"XPF": XPF,
	"XPT": XPT,
	"YER": YER,
	"ZAR": ZAR,
	"ZMW": ZMW,
	"ZWL": ZWL,
}

func NewGlobalCurrencyCode(code string) (CurrencyCode, error) {
	val, ok := currencyCodesRegister[code]
	if !ok {
		return CurrencyCode{}, fmt.Errorf("invalid or unsupported global currency code: %s", code)
	}

	return val, nil
}
