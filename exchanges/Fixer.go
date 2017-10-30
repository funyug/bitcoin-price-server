package exchanges

import (
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/funyug/bitcoin-price-server/controllers"
)

type Fixer struct {
	Rates Rate
}

type Rate struct {
	INR float64
}

func GetUSDRate(c *models.Client, price *controllers.BitcoinPrice) {
	rsp := &Fixer{}
	e := c.LoadResponse("GET","https://api.fixer.io/latest?base=USD",rsp)
	if(e != nil) {
		fmt.Print(e)
	} else {
		price.UsdRate = rsp.Rates.INR;
	}
}