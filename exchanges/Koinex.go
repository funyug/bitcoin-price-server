package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
	"strconv"
)

type Koinex struct {
	Stats struct {
		BTC struct {
			Lowest_ask string
			Highest_bid string
		}
	}
}

func GetKoinexPrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &Koinex{}
	e := c.LoadResponse("GET","https://koinex.in/api/ticker",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
			price.KoinexSellPrice,_ = strconv.ParseFloat(rsp.Stats.BTC.Highest_bid,64);
			price.KoinexBuyPrice,_ = strconv.ParseFloat(rsp.Stats.BTC.Lowest_ask,64);
			//controllers.SendExchangeAlerts(2,price.KoinexBuyPrice,price.KoinexSellPrice);
	}
}