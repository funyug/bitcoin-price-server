package controllers

import (
	"github.com/labstack/echo"
)

type BitcoinPrice struct {
	CoinSecureBuyPrice float64 `json:"coinsecureBuyPrice"`
	CoinSecureSellPrice float64 `json:"coinsecureSellPrice"`
	ZebpayBuyPrice float64 `json:"zebpayBuyPrice"`
	ZebpaySellPrice float64 `json:"zebpaySellPrice"`
	PocketBitsBuyPrice float64 `json:"pocketBitsBuyPrice"`
	PocketBitsSellPrice float64 `json:"pocketBitsSellPrice"`
	UsdRate float64 `json:"usd_rate"`
}


func GetBitcoinPrice(price *BitcoinPrice) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, price)
	}
}