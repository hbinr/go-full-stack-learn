package service

// IUserService 用户处理业务接口
type IUserService interface {
	GetUserName(userID int) string
}

type UserService struct {
}

func (u *UserService) GetUserName(userID int) string {
	if userID == 101 {
		return "101 test"
	}
	return "guest"
}
