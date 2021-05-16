package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

type PostgresDb struct {
	DB *gorm.DB
}

type DatabaseService interface {
	CreateUser(user *User) error
	DeleteUser(user *User) error
	UpdateUser(user *User) error
	GetAllUsers()([]User, error)
}

func SetDbEnv() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "crud")
	os.Setenv("DB_DBNAME", "postgres")
	os.Setenv("DB_PASSWORD", "1234")
	os.Setenv("DB_PORT", "5432")
}


func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s",os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_DBNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	return db
}

func (p PostgresDb) CreateUser(user *User) error {
	p.DB = ConnectDb()
	p.DB.Create(&User{Name: user.Name})
	return nil
}

func (p PostgresDb) DeleteUser(user *User) error {
	p.DB = ConnectDb()
	p.DB.Delete(&User{}, user.ID)
	return nil
}

func (p PostgresDb) UpdateUser(user *User) error {
	p.DB = ConnectDb()
	p.DB.Model(&User{}).Where("id = ?", user.ID).Update("Name", user.Name)
	return nil
}

func (p PostgresDb) GetAllUsers()([]User, error) {
	p.DB = ConnectDb()
	var user []User
	p.DB.Find(&user)
	return user, nil
}
