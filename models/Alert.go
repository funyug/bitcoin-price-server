package models

import (
	"github.com/jinzhu/gorm"
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