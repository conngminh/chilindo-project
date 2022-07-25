package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"practice/src/user-service/model"
)

var (
	username string = "root"
	password string = "Ronaldokl10112000"
	host     string = "localhost"
	port     string = "3306"
	database string = "practice"
)

var connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	username,
	password,
	host,
	port,
	database,
)

var Db *gorm.DB
var err error

func ConnectDatabase() {
	Db, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic("Connect database: error to connect database")
	}
	log.Println("Connect database:successfully")
	if err := Db.AutoMigrate(&model.User{}); err != nil {
		log.Println("Create table database: fail to migrate database!")
		return
	}
	log.Println("Create table database: migrate database successfully!")
}

func init() {
	ConnectDatabase()
}

func GetDB() *gorm.DB {
	return Db
}
