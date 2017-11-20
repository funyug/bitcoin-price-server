package controllers

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"github.com/funyug/bitcoin-price-server/models"
)

type BitcoinPrice struct {
	CoinSecureBuyPrice float64 `json:"coinsecureBuyPrice"`
	CoinSecureSellPrice float64 `json:"coinsecureSellPrice"`
	ZebpayBuyPrice float64 `json:"zebpayBuyPrice"`
	ZebpaySellPrice float64 `json:"zebpaySellPrice"`
	PocketBitsBuyPrice float64 `json:"pocketBitsBuyPrice"`
	PocketBitsSellPrice float64 `json:"pocketBitsSellPrice"`
	KoinexBuyPrice float64 `json:"koinexBuyPrice"`
	KoinexSellPrice float64 `json:"koinexSellPrice"`
	UsdRate float64 `json:"usd_rate"`
}


func GetBitcoinPrice(price *BitcoinPrice) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, price)
	}
}

func GetAlerts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		device_id := c.QueryParam("device_id")
		alerts, err := models.GetAlerts(db, device_id)
		if err != nil {
			response := models.Fail(err)
			return c.JSON(500,response)
		}
		response := models.Success(alerts)
		return c.JSON(200,response)
	}
}

func PostAlert(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//TODO: Add Validator
		alerts, err := models.PostAlert(db, c.Request())
		if err != nil {
			response := models.Fail(err)
			return c.JSON(500,response)
		}
		response := models.Success(alerts)
		return c.JSON(200,response)
	}
}

func DeleteAlert(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//TODO: Add Validator
		alerts, err := models.DeleteAlert(db, c.Request())
		if err != nil {
			response := models.Fail(err)
			return c.JSON(500,response)
		}
		response := models.Success(alerts)
		return c.JSON(200,response)
	}
}