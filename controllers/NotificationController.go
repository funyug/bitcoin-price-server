package controllers

import (
	"gopkg.in/maddevsio/fcm.v1"
	"fmt"
	"github.com/funyug/bitcoin-price-server/models"
	"strconv"
)

func SendNotification(token string,title string,body string) {
	data := map[string]string{}
	c := fcm.NewFCM("serverKey")
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: title,
			Body:  body,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
}

func SendAlerts(alerts []models.Alert, exchange int, price float64 ) {
	var exchange_name string;
	for i:=0;i<len(alerts);i++ {
		if exchange == 1 {
			exchange_name = "CoinSecure"
		} else if exchange == 2 {
			exchange_name = "Koinex"
		} else if exchange == 3 {
			exchange_name = "PocketBits"
		} else if exchange == 4 {
			exchange_name = "Zebpay"
		}
		SendNotification(alerts[i].DeviceId,"Price Alert",exchange_name + " Price crossed the thresold Price Rs " + strconv.FormatFloat(alerts[i].AlertPrice,'e',2,64));
	}
}

func SendExchangeAlerts(exchange int, buy_price float64, sell_price float64) {
	db, err := models.InitDB();
	if err != nil {
		panic(err)
	}
	defer db.Close()

	alerts := []models.Alert{};

	db.Where("exchange_id",exchange).Where("alert_price > ?",buy_price).Where("operator =",1).Where("price_type =",1).Find(&alerts);
	SendAlerts(alerts,exchange,buy_price)

	db.Where("exchange_id",exchange).Where("alert_price > ?",sell_price).Where("operator =",1).Where("price_type =",0).Find(&alerts);
	SendAlerts(alerts,exchange,sell_price)

	db.Where("exchange_id",exchange).Where("alert_price < ?",buy_price).Where("operator =",0).Where("price_type =",1).Find(&alerts);
	SendAlerts(alerts,exchange,buy_price)

	db.Where("exchange_id",exchange).Where("alert_price < ?",sell_price).Where("operator =",0).Where("price_type =",0).Find(&alerts);
	SendAlerts(alerts,exchange,buy_price)

}
