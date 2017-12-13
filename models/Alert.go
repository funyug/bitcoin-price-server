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
	PriceType int
	Operator int
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

	err := ValidatePostAlert(alert)
	if err == nil {
		FindOrCreate(db,alert.DeviceId)
		alert.Id = 0
		if alert.PriceType > 0 {
			alert.PriceType = 1
		}
		db.Create(&alert)
		alerts, err := GetAlerts(db,alert.DeviceId)
		return alerts, err
	}

	return []Alert{}, err
}

func DeleteAlert(db *gorm.DB, request *http.Request) ([]Alert, error) {
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	alert := Alert{}
	json.Unmarshal(b, &alert)

	err := ValidateDeleteAlert(alert)
	if err == nil {
		db.Where("id = ? and device_id = ?",alert.Id,alert.DeviceId).Delete(alert)
		alerts, err := GetAlerts(db,alert.DeviceId)
		return alerts, err
	}

	return []Alert{}, err
}

func ValidatePostAlert(alert Alert) error{
	if alert.DeviceId == "" {
		err := errors.New("Device Id not found")
		return err;
	}
	if alert.AlertPrice == 0 {
		err := errors.New("Alert Price not found")
		return err;
	}
	if alert.ExchangeId == 0 {
		err := errors.New("Exchange Id not found")
		return err;
	}
	return nil
}

func ValidateDeleteAlert(alert Alert) error{
	if alert.DeviceId == "" {
		err := errors.New("Device Id not found")
		return err;
	}
	if alert.Id == 0 {
		err := errors.New("Id not found")
		return err;
	}
	return nil
}