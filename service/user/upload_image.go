package user

import (
	"context"
	"mime/multipart"
)

func (s *UserService) UploadImage(ctx context.Context, file multipart.File, fileName string) error {
	err := s.aws.UploadImage(file, fileName)
	if err != nil {
		return err
	}
	return nil
}
