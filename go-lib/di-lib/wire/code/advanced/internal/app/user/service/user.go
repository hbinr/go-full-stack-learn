package service

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
	"github.com/google/wire"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/dao"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/model"
)

// 正常情况下会重新建一个UserDto结构体来做数据传输
type UserDto = model.User

// 验证接口是否实现
var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Insert(user *UserDto) error
	Delete(int64) bool
	Update(user *UserDto) error
	SelectById(id int64) (*UserDto, error)
	SlectByName(userName string) (*UserDto, error)
}

type UserService struct {
	Dao dao.IUserDao
}

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)))

//func NewUserService(db *gorm.DB) IUserService {
//	return &UserService{
//		Dao: dao.NewUserDao(db),
//	}
//}

func (r *UserService) Insert(user *UserDto) (err error) {
	var u dao.UserModel
	if err := gconv.Struct(user, &u); err != nil {
		fmt.Println("gconv.Struct(app, &u) failed,err:", err)
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
