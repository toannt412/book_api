package user

import (
	"bookstore/serialize"
	"context"
)

func (s *UserService) EditUser(ctx context.Context, id string, sth *serialize.User) (*serialize.User, error) {
	user, err := s.userRepo.EditUser(ctx, id, sth)
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

func (s *UserService) EditUserToken(ctx context.Context, token string) error {
	err := s.userRepo.EditUserToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
