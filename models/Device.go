package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	Id int
	Device_id string
}

func GetDevice(db *gorm.DB, device_id string) Device{
	device := Device{}
	db.Where("device_id = ?",device_id).First(&device)
	return device
}

func FindOrCreate(db *gorm.DB, device_id string) Device {
	device := GetDevice(db, device_id)

	if device.Id == 0 {
		device = Device{Device_id:device_id}
		db.Create(&device)
	}

	return device
}