package models

import (
	"fmt"

	"gorm.io/driver/mysql"

	"github.com/nazordz/boutique/product-catalog-service/configs"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.EnvConfigs.DBUser,
		configs.EnvConfigs.DBPassword,
		configs.EnvConfigs.DBHost,
		configs.EnvConfigs.DBPort,
		configs.EnvConfigs.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
