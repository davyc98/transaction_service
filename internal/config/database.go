package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Connect connection DB
func ConnectDB(host, database string, port int, username, password string) (db *gorm.DB, err error) {
	connectionString := fmt.Sprint("host=", host, " user=", username, " password=", password, " dbname=", database, " port=", port, " sslmode=disable TimeZone=Asia/Jakarta")
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	db.LogMode(true)
	log.Println("database connected")

	return db, err
}
