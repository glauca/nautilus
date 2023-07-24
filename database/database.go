package database

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	database := viper.GetString("DB_DATABASE")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	charset := viper.GetString("DB_CHARSET")
	prefix := viper.GetString("DB_PREFIX")

	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	DBConn = db
}
