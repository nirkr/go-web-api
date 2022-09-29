package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPGClient(dbUrl string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
