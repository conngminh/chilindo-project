package repository

import (
	"gorm.io/gorm"
	"log"
	"practice/src/user-service/model"
)

type IUserRepository interface {
	CreateAddress(address *model.Address) (*model.Address, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u UserRepository) CreateAddress(address *model.Address) (*model.Address, error) {
	record := u.Db.Create(&address)
	if record.Error != nil {
		log.Println("error to create address in package repository", record.Error)
		return nil, record.Error
	}
	return address, nil
}
