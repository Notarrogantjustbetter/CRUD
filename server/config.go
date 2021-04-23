package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	user = "crud"
	dbname = "postgres"
	password = 1234
	port = 5432
)

func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%d port=%d", host, user, dbname, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

