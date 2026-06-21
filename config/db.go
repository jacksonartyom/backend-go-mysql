package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gomysql "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := gomysql.Config{
		User:                 "root",
		Passwd:               "admin1234",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "testdb",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("connected database")
}
