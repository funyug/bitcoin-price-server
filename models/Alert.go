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
	Exchange_id int
	Alert_price float64
	Device_id string
}

func GetAlerts(db *gorm.DB, device_id string) ([]Alert, error) {
	alerts := []Alert{};
	device := GetDevice(db, device_id)
	if device.Id != 0 {
		db.Where("device_id = ?",device.Device_id).Find(&alerts);
		return alerts,nil
	}
	return []Alert{}, errors.New("Device not found")
}

func PostAlert(db *gorm.DB, request *http.Request) ([]Alert, error) {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	if alert.Device_id != "" {
		FindOrCreate(db,alert.Device_id)
		db.Create(&alert)
	}

	alerts, err := GetAlerts(db,alert.Device_id)
	return alerts, err
}

func DeleteAlert(db *gorm.DB, request *http.Request) ([]Alert, error) {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	if alert.Id != 0 && alert.Device_id != "" {
		db.Where("id = ? and device_id = ?",alert.Id,alert.Device_id).Delete(alert)
	}

	alerts, err := GetAlerts(db,alert.Device_id)
	return alerts, err
}