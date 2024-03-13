package mysql

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(c *config.Config) (*gorm.DB, error) {
	creds := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", c.MySQL.User, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.DBName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: creds,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Product{}, &models.Supplier{})

	return db, nil
}
