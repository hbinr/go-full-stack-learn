package service

// IUserService 用户处理业务接口
type IUserService interface {
	GetUserName(id int) string
}

type UserService struct {
}

func (u *UserService) GetUserName(id int) string {
	if id == 101 {
		return "101 test"
	}
	return "guest"
}
