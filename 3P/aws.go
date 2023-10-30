package third_party

import (
	"bookstore/configs"
	"context"
	"log"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSService struct {
	S3Client      *s3.Client
	PresignClient *s3.PresignClient
}

func NewAWSService() *AWSService {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configs.Config.AWSRegion))
	return &AWSService{
		S3Client: s3.NewFromConfig(cfg),
		PresignClient: s3.NewPresignClient(
			s3.NewFromConfig(cfg),
		),
	}
}

func (a *AWSService) UploadImage(file multipart.File, fileName string) error {
	_, err := a.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(configs.Config.BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		log.Println("Error uploading file", err)
	}
	return err
}

func (a *AWSService) GeneratePresignedURL(objectKey string) (*v4.PresignedHTTPRequest, error) {
	request, err := a.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(configs.Config.BucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(10 * time.Minute)
	})
	if err != nil {
		log.Println("Couldn't get presigned URL", err)
	}

	return request, err
}

func (a *AWSService) PutObjectUsePresignedURL(objectKey string) (*v4.PresignedHTTPRequest, error) {
	request, err := a.PresignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(configs.Config.BucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(10 * time.Minute)
	})
	if err != nil {
		log.Println("Couldn't get a presigned request to put", err)
	}
	return request, err
}
