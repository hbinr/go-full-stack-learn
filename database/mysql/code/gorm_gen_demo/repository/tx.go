package repository

import (
	"context"
	"time"

	"hb.study/database/mysql/code/gorm_gen_demo/data/model"
	"hb.study/database/mysql/code/gorm_gen_demo/data/query"
)

// DeleteAndUpdate 事务处理——闭包用法
func (u userRepo) DeleteAndUpdate(ctx context.Context) error {
	item := model.User{
		UserID:   1,
		UserName: "test_for_tx",
		Age:      50,
	}

	// 在事务中执行一组操作
	return u.sqlClient.Transaction(func(tx *query.Query) error {
		user := tx.User

		// DELETE FROM `user` WHERE `user`.`user_id` = 3214506793
		if _, err := user.WithContext(ctx).Where(tx.User.UserID.Eq(3214506793)).Delete(); err != nil {
			return err
		}
		// UPDATE `user` SET `user_name`='test_for_tx',`updated_at`='2021-12-20 13:41:35.464' WHERE `user`.`user_id` = 1
		_, err := user.WithContext(ctx).Select(user.UserName).Where(user.UserID.Eq(item.UserID)).Updates(item)
		if err != nil {
			return err
		}

		return nil
	})
}

// DeleteAndCreate 事务处理——手动开启事务
func (u userRepo) DeleteAndCreate(ctx context.Context) error {
	item := model.User{
		UserID:    3,
		UserName:  "test_for_DeleteAndCreate",
		Age:       50,
		Password:  "12345",
		Email:     "2ssss34@test.com",
		Phone:     "1233333332",
		CreatedAt: time.Now(),
	}
	tx := u.sqlClient.Begin()
	if err := tx.User.WithContext(ctx).Create(&item); err != nil {
		_ = tx.Rollback()
		return err
	}

	if _, err := tx.User.WithContext(ctx).Where(tx.User.UserID.Eq(2)).Delete(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
