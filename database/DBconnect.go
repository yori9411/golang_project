package database

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DBconnect *gorm.DB
var err error

func SQLDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:8890)/db?charset=utf8mb4&parseTime=True&loc=Local"
	//連接MySQL
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
