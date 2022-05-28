package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (db *gorm.DB) {

	dbConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	fmt.Println(getConnectionString())
	db, err := gorm.Open(mysql.Open(getConnectionString()), dbConfig)

	if err = db.AutoMigrate(
		&Balance{},
	); err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	return db

}

func getConnectionString() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2FTokyo",
		"sol",
		"sol",
		"127.0.0.1",
		3306,
		"solana",
	)
}
