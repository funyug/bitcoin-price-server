package main

import (
	"github.com/labstack/echo"
	"github.com/funyug/bitcoin-price-server/controllers"
	"github.com/funyug/bitcoin-price-server/exchanges"
	"github.com/funyug/bitcoin-price-server/models"
	"github.com/labstack/echo/middleware"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := models.InitDB();
	if err != nil {
		panic(err)
	}
	defer db.Close()

	price := controllers.BitcoinPrice{ };

	c,_ := models.New();

	getBitcoinPrices(c,&price,[]string{});
	go getPrices(c,&price,30,[]string{"Zebpay","PocketBits","Koinex"});
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
	e.GET("/alerts",controllers.GetAlerts(db));
	e.POST("/alerts",controllers.PostAlert(db));
	e.DELETE("/alerts",controllers.DeleteAlert(db));

	e.Logger.Fatal(e.Start(":3001"))

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
		exchanges.GetKoinexPrice(c,price);
		exchanges.GetUSDRate(c,price);
	} else {
		for _, v := range exchanges_arr {
			if v == "Coinsecure" {
				exchanges.GetCoinSecurePrice(c,price);
			} else if v == "PocketBits" {
				exchanges.GetPocketBitsPrice(c,price);
			} else if v == "Zebpay" {
				exchanges.GetZebpayPrice(c,price);
			} else if v == "Koinex" {
				exchanges.GetKoinexPrice(c,price);
			} else if v == "USDRate" {
				exchanges.GetUSDRate(c,price);
			}
		}
	}
}
