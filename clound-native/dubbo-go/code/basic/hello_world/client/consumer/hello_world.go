package consumer

type UserRequest struct {
	ID int64 `json:"id"`
}

type UserReply struct {
	Name string `json:"name"`
}

type UserProvider struct {
	GetUserName func(id string) (*UserReply, error)
	GetUser     func(req interface{}) (*UserReply, error)
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}
