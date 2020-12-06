package repository

import (
	"gorm.io/gorm"
	"hb.study/clound-native/go-kit/code/app_advanced/internal/user/model"
)

type UserRepository interface {
	Create(user *model.User) error
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &user{db: db}
}

func (u *user) Create(user *model.User) error {
	return u.db.Create(user).Error
}
