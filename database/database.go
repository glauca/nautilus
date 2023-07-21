package database

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	database := viper.GetString("db.database")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	charset := viper.GetString("db.charset")

	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)

	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("db connected")

	DBConn = db
}
