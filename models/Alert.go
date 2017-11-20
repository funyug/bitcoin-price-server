package models

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/pkg/errors"
)

type Alert struct {
	Id int
	ExchangeId int
	AlertPrice float64
	AlertType int
	DeviceId string
}

func GetAlerts(db *gorm.DB, device_id string) ([]Alert, error) {
	alerts := []Alert{};
	device := GetDevice(db, device_id)
	if device.Id != 0 {
		db.Where("device_id = ?",device.DeviceId).Find(&alerts);
		return alerts,nil
	}
	return []Alert{}, errors.New("Device not found")
}

func PostAlert(db *gorm.DB, request *http.Request) ([]Alert, error) {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	if alert.DeviceId != "" {
		FindOrCreate(db,alert.DeviceId)
		alert.Id = 0
		if alert.AlertType > 0 {
			alert.AlertType = 1
		}
		db.Create(&alert)
	}

	alerts, err := GetAlerts(db,alert.DeviceId)
	return alerts, err
}

func DeleteAlert(db *gorm.DB, request *http.Request) ([]Alert, error) {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	if alert.Id != 0 && alert.DeviceId != "" {
		db.Where("id = ? and device_id = ?",alert.Id,alert.DeviceId).Delete(alert)
	}

	alerts, err := GetAlerts(db,alert.DeviceId)
	return alerts, err
}