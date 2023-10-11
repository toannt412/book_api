package service

import (
	"context"
)

func (s *AdminService) DeleteAdmin(ctx context.Context, id string) (string, error) {
	res, err := s.adminRepo.DeleteAdmin(ctx, id)
	if err != nil {
		return "Deleted fail", err
	}
	return res, nil
}
