package user

import (
	"bookstore/dao/user/model"
	"bookstore/serialize"
	"context"
)

func (s *UserService) GetUserByID(ctx context.Context, id string) (*serialize.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &serialize.User{
		Id:          user.Id,
		FullName:    user.FullName,
		Location:    user.Location,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
	}, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByUserName(ctx context.Context, username string) (model.User, error) {
	user, err := s.userRepo.GetUserByUserName(ctx, username)
	if err != nil {
		return model.User{}, err
	}
	return user,nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	user, err := s.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (s *UserService) GetUserToken(ctx context.Context, token string) (model.User, error) {
	res, err := s.userRepo.GetUserToken(ctx, token)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
