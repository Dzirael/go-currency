package controllers

import (
	"log"
	"net/http"

	"github.com/Dzirael/go-curenncy/internal/pkg/exchange"
	http_err "github.com/Dzirael/go-curenncy/pkg/http-err"
	"github.com/gin-gonic/gin"
)

// GetRate godoc
// @Summary Отримання поточного курсу USD до UAH
// @Description Запит повертає курс згідно https://bank.gov.ua/
// @ID rate
// @Produce json
// @Success 200 {float} ExchangeRateResponse
// @Failure 500 {object} ErrorResponse
// @Router /exchange/rate [get]
func GetRate(c *gin.Context) {
	// Hardcode, left ability to use querry parametr
	rate, err := exchange.GetExchangeRate("USD")
	if err != nil {
		http_err.NewError(c, http.StatusInternalServerError, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, rate)
	}
}
