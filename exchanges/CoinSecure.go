package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
)

type CoinSecure struct {
	Message Message
}

type Message struct {
	Ask float64
	Bid float64
}

func GetCoinSecurePrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &CoinSecure{}
	e := c.LoadResponse("GET","https://api.coinsecure.in/v1/exchange/ticker",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.CoinSecureBuyPrice = rsp.Message.Ask/100;
		price.CoinSecureSellPrice = rsp.Message.Bid/100;
	}
}