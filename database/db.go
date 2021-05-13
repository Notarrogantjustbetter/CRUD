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

type User struct {
	gorm.Model
	Name string
	Email string
}


func OpenDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%d port=%d", host, user, dbname, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	return db
}

func CreateUser(name, email string) error {
	db := OpenDb()
	return db.Create(&User{Name: name, Email: email}).Error
}

func DeleteUser(id string) error {
	db := OpenDb()
	return db.Delete(&User{}, id).Error
}

func UpdateUser(id, key, value string) error {
	db := OpenDb()
	return db.Model(&User{}).Where("id = ?", id).Update(key, value).Error
}

func GetUsers() []User {
	db := OpenDb()
	var user []User
	db.Find(&user)
	return user
}
