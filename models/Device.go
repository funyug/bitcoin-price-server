package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	Id int
	Device_id string
}

func GetDevice(db *gorm.DB, device_id string) int{
	device := Device{}
	db.Where("device_id = ?",device_id).First(&device)
	return device.Id
}
