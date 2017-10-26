package main

import (
	"github.com/labstack/echo"
	"bitcoin-price-server/controllers"
	"bitcoin-price-server/exchanges"
	"bitcoin-price-server/models"
	"github.com/labstack/echo/middleware"
	"time"
)

func main() {
	price := controllers.BitcoinPrice{ };

	c,_ := models.New();

	getBitcoinPrices(c,&price,[]string{});
	go getPrices(c,&price,30,[]string{"Zebpay","PocketBits"});
	go getPrices(c,&price,60,[]string{"Coinsecure"});
	go getPrices(c,&price,600,[]string{"USDRate"});

	e := echo.New();
	e.Use(middleware.Logger());
	e.Use(middleware.Recover());
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}));
	e.GET("/bitcoin-price",controllers.GetBitcoinPrice(&price));
	e.Logger.Fatal(e.Start(":1323"))

}

func getPrices(c *models.Client, price *controllers.BitcoinPrice, interval time.Duration, exchanges []string){
	for range time.Tick(time.Second * interval){
		getBitcoinPrices(c,price, exchanges);
	}
}

func getBitcoinPrices(c *models.Client, price *controllers.BitcoinPrice, exchanges_arr []string) {
	if len(exchanges_arr) == 0 {
		exchanges.GetZebpayPrice(c,price);
		exchanges.GetCoinSecurePrice(c,price);
		exchanges.GetPocketBitsPrice(c,price);
		exchanges.GetUSDRate(c,price);
	} else {
		for _, v := range exchanges_arr {
			if v == "Coinsecure" {
				exchanges.GetCoinSecurePrice(c,price);
			} else if v == "PocketBits" {
				exchanges.GetPocketBitsPrice(c,price);
			} else if v == "Zebpay" {
				exchanges.GetZebpayPrice(c,price);
			} else if v == "USDRate" {
				exchanges.GetUSDRate(c,price);
			}
		}
	}
}
