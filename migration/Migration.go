package migration

import (
	"log"

	"github.com/irdaislakhuafa/go-graphql-jwt/config"
	"github.com/irdaislakhuafa/go-graphql-jwt/entity"
)

func EnableMigration(isEnable bool) {
	log.Println("Auto Migration status :", isEnable)
	if isEnable {
		db := config.GetDB()
		db.AutoMigrate(&entity.User{})
		log.Println("Success migration")
	}
}
