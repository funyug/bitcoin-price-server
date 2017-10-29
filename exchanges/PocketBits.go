package exchanges

import (
	"fmt"
	"bitcoin-price-server/models"
	"bitcoin-price-server/controllers"
)

type PocketBits struct {
	BTC_SellingRate float64
	BTC_BuyingRate float64
}

func GetPocketBitsPrice(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &PocketBits{}
	e := c.LoadResponse("GET","https://www.pocketbits.co.in/Index/getBTCRate",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.PocketBitsBuyPrice = rsp.BTC_BuyingRate;
		price.PocketBitsSellPrice = rsp.BTC_SellingRate;
	}
}