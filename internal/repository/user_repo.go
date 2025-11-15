package repository

import (
	"runsystem-test/internal/entity"
	"runsystem-test/internal/helper"

	"gorm.io/gorm"
)

type UserRepo interface {
	SaveUser(*entity.UserEntity) (*entity.UserEntity, helper.ErrorInterface)
	GetUserByID(id int) (*entity.UserEntity, helper.ErrorInterface)
	GetUserByName(name string) (*entity.UserEntity, helper.ErrorInterface)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) SaveUser(user *entity.UserEntity) (*entity.UserEntity, helper.ErrorInterface) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, helper.NewInternalServerError("something went wrong")
	}
	return user, nil
}

func (r *userRepo) GetUserByID(id int) (*entity.UserEntity, helper.ErrorInterface) {
	var user entity.UserEntity
	result := r.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, helper.NewNotFoundError("user not found")
		}
		return nil, helper.NewInternalServerError("something went wrong")
	}
	return &user, nil
}

func (r *userRepo) GetUserByName(name string) (*entity.UserEntity, helper.ErrorInterface) {
	var user entity.UserEntity
	result := r.db.Where("name = ?", name).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, helper.NewNotFoundError("user not found")
		}
		return nil, helper.NewInternalServerError("something went wrong")
	}
	return &user, nil
}