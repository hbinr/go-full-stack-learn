package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"hb.study/mysql/code/gorm_gen_demo/data/model"
)

func (u userRepo) CreateUser(ctx context.Context) {
	user := u.sqlClient.User.WithContext(ctx)
	// 创建用户
	item := model.User{
		UserID:    1,
		Age:       2,
		UserName:  "test",
		Password:  "test",
		Email:     "test@test.com",
		Phone:     "12334569870",
		RoleName:  "member",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := user.Create(&item)
	if err != nil {
		return
	}
}

//  GetUser GROM 提供了 First、Take、Last 方法从数据库中查询单条数据，在查询数据库时会自动添加 LIMIT 1 条件，如果没有找到记录则返回错误 ErrRecordNotFound。
func (u userRepo) GetUser(ctx context.Context) (*model.User, error) {
	user := u.sqlClient.User.WithContext(ctx)

	// SELECT * FROM users ORDER BY id LIMIT 1;
	res, err := user.First()

	// SELECT * FROM users LIMIT 1;
	res, err = user.Take()

	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	res, err = user.Last()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, "record not found")
		}
		return nil, err
	}

	return res, err
}

func (u userRepo) GetUserByCondition(ctx context.Context) (*model.User, error) {
	user := u.sqlClient.User
	// SELECT * FROM users WHERE id = 6 LIMIT 1;
	res, err := user.WithContext(ctx).Where(user.ID.Eq(6)).First()
	fmt.Println("res", res)
	if err != nil {
		return nil, err
	}
	// SELECT * FROM `user` WHERE `user`.`id` > 0
	resArr, _ := user.WithContext(ctx).Where(user.ID.Gt(0)).Find()
	return resArr[0], nil
}

func (u userRepo) UpdateUser(ctx context.Context) {
	user := u.sqlClient.User.WithContext(ctx)
	fmt.Println("user", user)
}
