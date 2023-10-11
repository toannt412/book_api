package service

import (
	"bookstore/dao/admin/model"
	"bookstore/serialize"
	"context"
)

func (s *AdminService) GetAdminByID(ctx context.Context, id string) (*serialize.Admin, error) {

	user, err := s.adminRepo.GetAdminByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &serialize.Admin{
		Id:       user.Id.Hex(),
		UserName: user.UserName,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *AdminService) GetAdminByEmail(ctx context.Context, email string) error {

	_, err := s.adminRepo.GetAdminByEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) GetAdminByUserName(ctx context.Context, username string) error {
	_, err := s.adminRepo.GetAdminByUserName(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) GetAllAdmins(ctx context.Context) ([]model.Admin, error) {
	admins, err := s.adminRepo.GetAllAdmins(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (s *AdminService) GetAdminToken(ctx context.Context, token string) (model.Admin, error) {
	admin, err := s.adminRepo.GetAdminToken(ctx, token)
	if err != nil {
		return model.Admin{}, err
	}
	return admin, nil
}
