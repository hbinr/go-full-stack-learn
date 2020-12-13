package repository

import (
	"gorm.io/gorm"
	"hb.study/clound-native/go-kit/code/app_advanced/internal/user/model"
)

type IUserRepository interface {
	Create(user *model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) Create(user *model.User) error {
	return u.db.Create(user).Error
}
