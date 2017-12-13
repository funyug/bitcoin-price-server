package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
	"strconv"
)

type Koinex struct {
	Buy_orders struct {
		Data []struct {
			Price_per_unit string
		}
	}
	Sell_orders struct {
		Data []struct {
			Price_per_unit string
		}
	}
}

func GetKoinexPrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &Koinex{}
	e := c.LoadResponse("GET","https://koinex.in/api/dashboards/order_history?page=1&per_page=13&target_currency=bitcoin",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		if(len(rsp.Buy_orders.Data) > 0 && len(rsp.Sell_orders.Data) > 0) {
			price.KoinexSellPrice,_ = strconv.ParseFloat(rsp.Buy_orders.Data[0].Price_per_unit,64);
			price.KoinexBuyPrice,_ = strconv.ParseFloat(rsp.Sell_orders.Data[0].Price_per_unit,64);
			controllers.SendExchangeAlerts(2,price.KoinexBuyPrice,price.KoinexSellPrice);
		}
	}
}