package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DbCon *gorm.DB
}

func (db *Database) GetDB() *gorm.DB {
	return db.DbCon
}

func (db *Database) Init() {
	var err error
	log.Println("initilaizing database connection")
	const (
		username string = "root"
		password string = "p"
		hostname string = "localhost"
		port     string = "3306"
		database string = "test"
	)
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", username, password, hostname, port, database)

	log.Println("Open database connection to \"" + database + "\"")
	db.DbCon, err = gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error when open database connection: " + err.Error())
	} else {
		log.Println("Success connected to database")
	}
}
