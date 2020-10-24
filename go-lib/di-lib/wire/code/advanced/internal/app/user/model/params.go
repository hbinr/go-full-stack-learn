package model

// SignUpInput 注册入参
type SignUpInput struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Nickname string `json:"nickname" form:"nickname"`
}
