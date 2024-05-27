package user

import (
	"context"
)

func (s *UserService) DeleteUser(ctx context.Context, id string) (string, error) {
	res, err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return "Deleted fail", err
	}
	return res, nil
}
