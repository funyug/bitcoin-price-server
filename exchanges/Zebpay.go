package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
	"strconv"
)

type Zebpay struct {
	Buy string `json:"buy"`
	Sell string `json:"sell"`
	Market string `json:"market"`
	Currency string `json:"currency"`
	Volume float64 `json:"volume"`
}

func GetZebpayPrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &Zebpay{}
	e := c.LoadResponse("GET","https://www.zebapi.com/api/v1/market/ticker-new/BTC/INR",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.ZebpayBuyPrice,_ = strconv.ParseFloat(rsp.Sell,64);
		price.ZebpaySellPrice,_ = strconv.ParseFloat(rsp.Buy,64);
		controllers.SendExchangeAlerts(4,price.ZebpayBuyPrice,price.ZebpaySellPrice);
	}
}