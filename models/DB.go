package models

import "github.com/jinzhu/gorm"

func InitDB(db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:@/bitcoin_price")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
