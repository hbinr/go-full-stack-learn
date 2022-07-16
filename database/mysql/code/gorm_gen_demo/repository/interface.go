package repository

import (
	"context"

	"gorm.io/gorm"
	"hb.study/mysql/code/gorm_gen_demo/data/model"
	"hb.study/mysql/code/gorm_gen_demo/data/query"
)

type UserRepo interface {
	CreateUser(context.Context)

	// 查询相关
	GetUser(context.Context) (*model.User, error)
	GetUserByCondition(context.Context) (*model.User, error)
	GetUserBySelectFiled(context.Context) (*model.User, error)
	GetUserByPage(context.Context) ([]*model.User, error)
	GetSingleFiledByPluck(context.Context) ([]string, error)

	// 更新相关
	UpdateUser(context.Context) error
	UpdateSingleFiled(context.Context) error
	UpdateSelectFiled(context.Context) error
	UpdateOmitFiled(context.Context) error

	// 事务相关
	DeleteAndUpdate(context.Context) error
	DeleteAndCreate(context.Context) error
}

type userRepo struct {
	sqlClient *query.Query
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		sqlClient: query.Use(db),
	}
}
