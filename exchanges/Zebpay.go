package exchanges

import (
	"fmt"
	"bitcoin-price-server/models"
	"bitcoin-price-server/controllers"
)

type Zebpay struct {
	Buy float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Market float64 `json:"market"`
	Currency string `json:"currency"`
	Volume float64 `json:"volume"`
}

func GetZebpayPrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &Zebpay{}
	e := c.LoadResponse("https://api.zebpay.com/api/v1/ticker?currencyCode=INR",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.ZebpayBuyPrice = rsp.Buy;
		price.ZebpaySellPrice = rsp.Sell;
	}
}