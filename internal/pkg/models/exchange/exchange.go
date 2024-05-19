package exchange

type RateResponce []struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	Exchangedate string  `json:"exchangedate"`
}
