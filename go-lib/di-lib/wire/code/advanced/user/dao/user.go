package dao

import (
	"errors"
	"fmt"

	"github.com/google/wire"
	"gorm.io/gorm"

	"hb.study/go-lib/di-lib/wire/code/advanced/user/model"
)

// 验证接口是否实现
var _ UserDaoProvider = (*UserDao)(nil)

type UserModel = model.User

type UserDaoProvider interface {
	Insert(user *UserModel) error
	Delete(int64) bool
	Update(user *UserModel) error
	SelectById(id int64) (*UserModel, error)
	SlectByName(userName string) (*UserModel, error)
}

type UserDao struct {
	DB *gorm.DB
}

// NewUserDaoSet 使用wire 依赖注入，相当于下面的 NewUserDao 函数
var NewUserDaoSet = wire.NewSet(
	wire.Struct(new(UserDao), "*"),
	wire.Bind(new(UserDaoProvider), new(*UserDao)))

//func NewUserDao(db *gorm.DB) UserDaoProvider {
//	return &UserDao{
//		DB: db,
//	}
//}

func (r *UserDao) Insert(user *UserModel) (err error) {
	return r.DB.Create(&user).Error
}

func (r *UserDao) Delete(id int64) bool {
	return r.DB.Where("id = ?", id).Delete(&UserModel{}).Error == nil
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
