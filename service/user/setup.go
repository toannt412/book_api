package user

import "bookstore/dao/user"

type UserService struct {
	userRepo *user.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: user.NewUserRepository(),
	}
}
