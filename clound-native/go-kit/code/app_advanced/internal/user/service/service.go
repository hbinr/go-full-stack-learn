package service

import "hb.study/clound-native/go-kit/code/app_advanced/internal/user/repository"

type IUserService interface {
	Register() error
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo: repo}
}

func (u *userService) Register() error {
	panic("implement me")
}
