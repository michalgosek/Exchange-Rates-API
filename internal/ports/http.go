package ports

import (
	"exchange-rates-api/internal/app"
	"exchange-rates-api/internal/app/query"
	"exchange-rates-api/internal/core"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HTTP struct {
	app *app.Application
	eng *gin.Engine
}

func (h *HTTP) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.eng.ServeHTTP(w, req)
}

func (h *HTTP) GetCryptoExchangeRate(c *gin.Context, params GetCryptoExchangeRateParams) {
	query, err := query.NewCryptoExchangeRateQuery(params.From, params.To, params.Amount)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	exchange, err := h.app.Queries.CryptoExchangeRateHandler.Handle(c.Request.Context(), query)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, ConvertToExchangeRateDTO(exchange))
}

func (h *HTTP) GetGlobalExchangeRates(c *gin.Context, params GetGlobalExchangeRatesParams) {
	query, err := query.NewGlobalExchangeRatesQuery(params.Currencies, core.USD.String())
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	exchanges, err := h.app.Queries.GlobalExchangeRatesHandler.Handle(c.Request.Context(), query)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, ConvertToExchangeRateDTOs(exchanges...))
}

func NewHTTP(app *app.Application) *HTTP {
	if app == nil {
		panic("application component is required")
	}

	eng := gin.Default()
	eng.Use(cors.Default())
	http := HTTP{
		app: app,
		eng: eng,
	}

	RegisterHandlers(eng, &http)
	return &http
}

func ConvertToExchangeRateDTO(rate core.CalculatedExchangeRate) ExchangeRateDTO {
	return ExchangeRateDTO{
		From: rate.From(),
		Rate: rate.ExchangeRate().String(),
		To:   rate.To(),
	}
}

func ConvertToExchangeRateDTOs(rates ...core.CalculatedExchangeRate) []ExchangeRateDTO {
	var dtos []ExchangeRateDTO
	for _, r := range rates {
		dtos = append(dtos, ExchangeRateDTO{
			From: r.From(),
			To:   r.To(),
			Rate: r.ExchangeRate().String(),
		})
	}
	return dtos
}
