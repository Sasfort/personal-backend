package database

import (
	"log"

	"github.com/Sasfort/personal-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dns := "root:admin@tcp(127.0.0.1:3306)/personal?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.Character{}, &models.Origin{})

	Database = DbInstance{Db: db}
}
