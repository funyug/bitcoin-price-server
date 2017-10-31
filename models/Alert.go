package models

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Alert struct {
	Id int
	Exchange_id int
	Alert_price float64
	Device_id string
}

func GetAlerts(db *gorm.DB, device_id string) []Alert {
	alerts := []Alert{};
	id := GetDevice(db, device_id)
	db.Where("device_id = ?",id).Find(&alerts);
	return alerts
}

func PostAlert(db *gorm.DB, request *http.Request) []Alert {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	if alert.Device_id != "" {
		db.Create(&alert)
	}

	alerts := GetAlerts(db,alert.Device_id)
	return alerts
}