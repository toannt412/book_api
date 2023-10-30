package user

import v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"

func (s *UserService) UploadImageUsePresignedURL(objectKey string) (*v4.PresignedHTTPRequest, error) {
	request, err := s.aws.PutObjectUsePresignedURL(objectKey)
	if err != nil {
		return nil, err
	}
	return request, nil
}
