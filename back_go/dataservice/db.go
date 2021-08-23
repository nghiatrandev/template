package dataservice

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
}

func NewDB() (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}
