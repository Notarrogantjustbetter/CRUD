package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
}

type DatabaseService interface {
	CreateUser(name, email string)
	DeleteUser(id string)
	GetUsers() []User
	UpdateUser(id, key, value string)
}

func InitDb() {
	dsn := "host=localhost user=postgres dbname=postgres password=1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}

func (u User) CreateUser(name, email string) {
	dsn := "host=localhost user=postgres dbname=postgres password=1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	u.Name = name
	u.Email = email
	db.Create(&User{Name: u.Name, Email: u.Email})
}

func (u User) DeleteUser(id string) {
	dsn := "host=localhost user=postgres dbname=postgres password=1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Delete(&User{}, id)
}

func (u User) GetUsers() []User {
	dsn := "host=localhost user=postgres dbname=postgres password=1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var user []User
	db.Find(&user)
	return user
}

func (u User) UpdateUser(id, key, value string) {
	dsn := "host=localhost user=postgres dbname=postgres password=1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Model(&User{}).Where("id = ?", id).Update(key, value)
}
