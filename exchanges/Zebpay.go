package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
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
	e := c.LoadResponse("GET","https://www.zebapi.com/api/v1/market/ticker-new/BTC/INR",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.ZebpayBuyPrice = rsp.Sell;
		price.ZebpaySellPrice = rsp.Buy;
		controllers.SendExchangeAlerts(4,price.ZebpayBuyPrice,price.ZebpaySellPrice);
	}
}