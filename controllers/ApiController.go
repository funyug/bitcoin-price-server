package controllers

import (
	"github.com/labstack/echo"
)

type BitcoinPrice struct {
	CoinSecureBuyPrice float64
	CoinSecureSellPrice float64``
	ZebpayBuyPrice float64
	ZebpaySellPrice float64
	PocketBitsBuyPrice float64
	PocketBitsSellPrice float64
}


func GetBitcoinPrice(price *BitcoinPrice) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, price)
	}
}