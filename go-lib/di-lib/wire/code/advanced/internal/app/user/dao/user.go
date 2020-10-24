package dao

import (
	"errors"
	"fmt"

	"github.com/google/wire"
	"gorm.io/gorm"

	"hb.study/go-lib/di-lib/wire/code/advanced/user/model"
)

// 验证接口是否实现
var _ IUserDao = (*UserDao)(nil)

type UserModel = model.User

type IUserDao interface {
	Insert(user *UserModel) error
	Delete(int64) bool
	Update(user *UserModel) error
	SelectById(id int64) (*UserModel, error)
	SlectByName(userName string) (*UserModel, error)
}

type UserDao struct {
	DB *gorm.DB
}

// UserDaoSet 使用wire 依赖注入，相当于下面的 NewUserDao 函数
var UserDaoSet = wire.NewSet(
	wire.Struct(new(UserDao), "*"),
	wire.Bind(new(IUserDao), new(*UserDao)))

//func NewUserDao(db *gorm.DB) IUserDao {
//	return &UserDao{
//		DB: db,
//	}
//}

func (r *UserDao) Insert(user *UserModel) (err error) {
	return r.DB.Create(&user).Error
}

func (r *UserDao) Delete(id int64) bool {
	// 这种删除调用sql: UPDATE `user` SET `deleted_at`='2020-10-24 23:54:36.003' WHERE id = 6
	// 重复请求删除还是会再一次设置时间，即使之前已经删除过了。因为有deleted_at字段。可以不用这个字段
	//return r.DB.Debug().Where("id = ?", id).Delete(&UserModel{}).Error == nil

	return r.DB.Debug().Where("id = ?", id).Delete(&UserModel{}).RowsAffected > 0
}

func (r *UserDao) Update(user *UserModel) error {
	return r.DB.Where("id = ?").Updates(&user).Error
}

func (r *UserDao) SlectByName(userName string) (*UserModel, error) {
	var user UserModel
	if err := r.DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) SelectById(id int64) (*UserModel, error) {
	var (
		user UserModel
		err  error
	)
	if err = r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("gorm.ErrRecordNotFound", err)
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}
