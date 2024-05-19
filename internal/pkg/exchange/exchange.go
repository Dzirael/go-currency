package exchange

import (
	"encoding/json"
	"errors"
	"time"

	models "github.com/Dzirael/go-curenncy/internal/pkg/models/exchange"
	"github.com/go-resty/resty/v2"
)

func GetExchangeRate(currencyTicker string) (float64, error) {
	client := resty.New().SetBaseURL("https://bank.gov.ua/NBUStatService/v1/statdirectory")

	params := map[string]string{
		"valcode": currencyTicker,
		"date":    time.Now().Format("20060102"),
		"json":    "",
	}

	resp, err := client.R().
		SetQueryParams(params).
		Get("/exchange")

	if err != nil {
		return 0, err
	}

	var rateResp models.RateResponce
	if err := json.Unmarshal(resp.Body(), &rateResp); err != nil {
		return 0, err
	}

	if len(rateResp) != 0 {
		return rateResp[0].Rate, nil
	} else {
		return 0, errors.New("rate not recieved")
	}
}
