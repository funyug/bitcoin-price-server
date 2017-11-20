package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	Id int
	DeviceId string
}

func GetDevice(db *gorm.DB, device_id string) Device{
	device := Device{}
	db.Where("device_id = ?",device_id).First(&device)
	return device
}

func FindOrCreate(db *gorm.DB, device_id string) Device {
	device := GetDevice(db, device_id)

	if device.Id == 0 {
		device = Device{DeviceId:device_id}
		db.Create(&device)
	}

	return device
}