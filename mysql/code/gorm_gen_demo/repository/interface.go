package repository

import (
	"context"

	"gorm.io/gorm"
	"hb.study/mysql/code/gorm_gen_demo/data/model"
	"hb.study/mysql/code/gorm_gen_demo/data/query"
)

type UserRepo interface {
	CreateUser(context.Context)
	GetUser(context.Context) (*model.User, error)
	GetUserByCondition(context.Context) (*model.User, error)
	UpdateUser(context.Context)
}

type userRepo struct {
	sqlClient *query.Query
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		sqlClient: query.Use(db),
	}
}
