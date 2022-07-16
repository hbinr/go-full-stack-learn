package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"hb.study/database/mysql/code/gorm_gen_demo/data/model"
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
	res, _ := user.WithContext(ctx).Where(user.ID.Eq(6)).First()
	fmt.Println("res", res)

	// SELECT * FROM `user` WHERE `user`.`id` > 0
	resArr, _ := user.WithContext(ctx).Where(user.ID.Gt(0)).Find()
	return resArr[0], nil
}

func (u userRepo) GetUserBySelectFiled(ctx context.Context) (*model.User, error) {
	user := u.sqlClient.User
	//  SELECT user_name, age FROM `user` ORDER BY id LIMIT 1;
	res, _ := user.WithContext(ctx).Select(user.UserName, user.Age).First()

	//  SELECT `user`.`user_name`,`user`.`age` FROM `user` WHERE `user`.`id` = 1 OR `user`.`age` > 10 ORDER BY `user`.`id` LIMIT 1
	res, _ = user.WithContext(ctx).Select(user.UserName, user.Age).Where(user.ID.Eq(1)).Or(user.Age.Gt(10)).First()
	return res, nil
}

func (u userRepo) GetUserByPage(ctx context.Context) ([]*model.User, error) {
	user := u.sqlClient.User
	// 一页5条, 第二页; offset 一定是5的倍数, 第一页是从0开始
	//  SELECT * FROM `user` limit 5 offset 5
	res, _ := user.WithContext(ctx).Limit(5).Offset(5).Find()

	// 第一页
	//  SELECT * FROM `user` limit 5
	res, _ = user.WithContext(ctx).Limit(5).Offset(0).Find()
	return res, nil
}

//  GetSingleFiledByPluck Pluck 方法支持从数据库中查询单列并扫描成切片。如果要查询多列，请使用 Select 和 Scan 方法代替 Pluck 方法
func (u userRepo) GetSingleFiledByPluck(ctx context.Context) ([]string, error) {
	user := u.sqlClient.User

	var names []string
	// SELECT `user_name` FROM `user` WHERE `user`.`age` > 5
	user.WithContext(ctx).Where(user.Age.Gt(5)).Pluck(user.UserName, &names)

	return names, nil
}

func (u userRepo) UpdateUser(ctx context.Context) error {
	user := u.sqlClient.User
	// 修改用户
	item := model.User{
		UserID:    1,
		Age:       2,
		UserName:  "test_update",
		Password:  "test_update",
		Email:     "test_update@test.com",
		Phone:     "12334569870",
		RoleName:  "member",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// 更新所有非空字段 where id = ?
	_, err := user.WithContext(ctx).Where(user.ID.Eq(item.ID)).Updates(item)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepo) UpdateSingleFiled(ctx context.Context) error {
	user := u.sqlClient.User
	// 修改用户
	item := model.User{
		UserID: 1,
	}
	// 更新单个字段
	// UPDATE `user` SET `age`=15,`updated_at`='2021-12-17 17:22:13.977' WHERE `user`.`user_id` = 1
	_, err := user.WithContext(ctx).Where(user.UserID.Eq(item.UserID)).Update(user.Age, 15)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepo) UpdateSelectFiled(ctx context.Context) error {
	user := u.sqlClient.User
	// 修改用户
	item := model.User{
		UserID:   1,
		UserName: "test_update_select",
		Age:      50,
	}
	// 更新指定字段 可以使用 Select
	// UPDATE `user` SET `user_name`='test_update_select',`updated_at`='2021-12-17 17:29:21.639' WHERE `user`.`user_id` = 1
	_, err := user.WithContext(ctx).Select(user.UserName).Where(user.UserID.Eq(item.UserID)).Updates(item)
	if err != nil {
		return err
	}
	// 更新时忽略某些字段 可以使用 Omit
	return nil
}

func (u userRepo) UpdateOmitFiled(ctx context.Context) error {
	user := u.sqlClient.User
	// 修改用户
	item := model.User{
		UserID:   1,
		UserName: "test_update_select",
		Age:      50,
	}
	// 更新时忽略某些字段 可以使用 Omit
	// UPDATE `user` SET `age`=50,`updated_at`='2021-12-17 17:34:23.576' WHERE `user`.`user_id` = 1
	_, err := user.WithContext(ctx).Omit(user.UserName, user.UserID).Where(user.UserID.Eq(item.UserID)).Updates(item)
	if err != nil {
		return err
	}

	return nil
}
