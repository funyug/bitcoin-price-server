package main

import (
	"github.com/labstack/echo"
	"bitcoin-price-server/controllers"
	"bitcoin-price-server/exchanges"
	"bitcoin-price-server/models"
	"time"
)

func main() {
	price := controllers.BitcoinPrice{ };

	c,_ := models.New();

	getBitcoinPrices(c,&price);
	go getPrices(c,&price)

	e := echo.New()
	e.GET("/bitcoin-price",controllers.GetBitcoinPrice(&price));
	e.Logger.Fatal(e.Start(":1323"))

}

func getPrices(c *models.Client, price *controllers.BitcoinPrice){
	for range time.Tick(time.Second *30){
		getBitcoinPrices(c,price);
	}
}

func getBitcoinPrices(c *models.Client, price *controllers.BitcoinPrice) {
	exchanges.GetZebpayPrice(c,price);
	exchanges.GetCoinSecurePrice(c,price);
	exchanges.GetPocketBitsPrice(c,price);
	exchanges.GetUSDRate(c,price);
}
