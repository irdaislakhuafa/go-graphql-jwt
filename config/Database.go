package config

import "gorm.io/gorm"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
