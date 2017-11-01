package models

import "github.com/jinzhu/gorm"

func InitDB() (*gorm.DB,error) {
	db, err := gorm.Open("mysql", "root:@/bitcoin_price")
	return db, err
}
