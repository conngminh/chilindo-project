package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"practice/src/product-service/models"
)

var Instance *gorm.DB
var dbError error

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

func ConnectDatabase() {
	//Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	//if dbError != nil {
	//	panic("Connect: Error connect to DB")
	//}
	//log.Println("Connected to Database!")
	Instance, dbError = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if dbError != nil {
		panic("Connect database: error to connect database")
	}
	log.Println("Connect database:successfully")
	if err := Instance.AutoMigrate(&models.Product{}, &models.Image{}, &models.Option{}); err != nil {
		log.Println("Create table database: fail to migrate database!", err)
		return
	}
	log.Println("Create table database: migrate database successfully!")
}

func init() {
	ConnectDatabase()
}

func GetDB() *gorm.DB {
	return Instance
}
