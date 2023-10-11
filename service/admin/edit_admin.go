package service

import (
	"bookstore/serialize"
	"context"
)

func (s *AdminService) EditAdmin(ctx context.Context, id string, sth *serialize.Admin) (*serialize.Admin, error) {
	user, err := s.adminRepo.EditAdmin(ctx, id, sth)
	if err != nil {
		return nil, err
	}
	return &serialize.Admin{
		FullName: user.FullName,
		Phone:    user.Phone,
		Role:     user.Role,
	}, nil
}

func (s *AdminService) EditAdminToken(ctx context.Context, token string) error {
	err := s.adminRepo.EditAminToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
