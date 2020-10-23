package service

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
	"github.com/google/wire"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/dao"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/model"
)

type UserDto = model.User

// 验证接口是否实现
var _ UserServiceProvider = (*UserService)(nil)

type UserServiceProvider interface {
	Insert(user *UserDto) error
	Delete(int64) bool
	Update(user *UserDto) error
	SelectById(id int64) (*UserDto, error)
	SlectByName(userName string) (*UserDto, error)
}

type UserService struct {
	Dao dao.UserDaoProvider
}

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(UserServiceProvider), new(*UserService)))

//func NewUserService(db *gorm.DB) UserServiceProvider {
//	return &UserService{
//		Dao: dao.NewUserDao(db),
//	}
//}

func (r *UserService) Insert(user *UserDto) (err error) {
	var u dao.UserModel
	if err := gconv.Struct(user, &u); err != nil {
		fmt.Println("gconv.Struct(user, &u) failed,err:", err)
		return err
	}
	return r.Dao.Insert(&u)
}

func (r *UserService) Delete(id int64) bool {
	return r.Dao.Delete(id)
}

func (r *UserService) Update(user *UserDto) error {
	return r.Dao.Update(user)
}

func (r *UserService) SlectByName(userName string) (*UserDto, error) {
	return r.Dao.SlectByName(userName)
}

func (r *UserService) SelectById(id int64) (*UserDto, error) {
	return r.Dao.SelectById(id)
}
