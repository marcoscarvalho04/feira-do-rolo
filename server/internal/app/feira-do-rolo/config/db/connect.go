package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openConnectionToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password=postgres dbname=feiradorolo port=5432 TimeZone=America/Sao_Paulo",
	}), &gorm.Config{})

	if err != nil {
		return &gorm.DB{}, err
	}
	return db, nil
}
