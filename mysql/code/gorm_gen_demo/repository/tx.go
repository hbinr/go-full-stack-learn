package repository

import (
	"context"

	"hb.study/mysql/code/gorm_gen_demo/data/query"
)

// 事务处理

func (u userRepo) DeleteAndUpdate(ctx context.Context) error {
	// item := model.User{
	// 	UserID:   1,
	// 	UserName: "test_update_select",
	// 	Age:      50,
	// }

	// 在事务中执行一组操作
	u.sqlClient.Transaction(func(tx *query.Query) error {
		// 删除
		if _, err := tx.User.WithContext(ctx).Where(tx.User.UserID.Eq(1)).Delete(); err != nil {
			return err
		}

		return nil
	})

	// 更新时忽略某些字段 可以使用 Omit
	// UPDATE `user` SET `age`=50,`updated_at`='2021-12-17 17:34:23.576' WHERE `user`.`user_id` = 1
	// _, err := user.WithContext(ctx).Omit(user.UserName, user.UserID).Where(user.UserID.Eq(item.UserID)).Updates(item)
	// if err != nil {
	// 	return err
	// }

	return nil
}
