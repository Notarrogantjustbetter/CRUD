package database

import "gorm.io/gorm"

type PostgresDB struct {}

type User struct {
	gorm.Model
	Name string
	Email string
}

type DatabaseService interface {
	CreateSchema()
	CreateUser(name, email string) (*User, error)
	DeleteUser(id string) (*User, error)
	UpdateUser(id, key, value string) (*User, error)
	GetUsers() []User
}

func (p PostgresDB) CreateSchema() {
	connDb := ConnectDb()
	connDb.AutoMigrate(&User{})
}

func (p PostgresDB) CreateUser(name, email string) (*User, error) {
	db := ConnectDb()
	db.Create(&User{Name: name, Email: email})
	return &User{}, nil
}

func (p PostgresDB) DeleteUser(id string) (*User, error) {
	db := ConnectDb()
	db.Delete(&User{}, id)
	return &User{}, nil
}

func (p PostgresDB) UpdateUser(id, key, value string) (*User, error) {
	db := ConnectDb()
	db.Model(&User{}).Where("id = ?", id).Update(key, value)
	return &User{}, nil
}

func (p PostgresDB) GetUsers() []User {
	db := ConnectDb()
	var user []User
	db.Find(&user)
	return user
}
