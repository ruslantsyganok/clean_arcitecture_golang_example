package service

import (
	"bytes"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"zen_api/internal/repository"
)

type FileUploaderService interface {
	UploadFile(file []byte) (string, error)
	DeleteFile(filePath string) error
}

type fileUploaderService struct {
	dao        repository.DAO
	bucketName string
	host       string
	keyID      string
	keySecret  string
	region     string
}

func NewFileUploaderService(
	dao repository.DAO,
	bucketName,
	host,
	region,
	keyID,
	keySecret string) FileUploaderService {
	return &fileUploaderService{
		dao:        dao,
		bucketName: bucketName,
		host:       host,
		region:     region,
		keyID:      keyID,
		keySecret:  keySecret}
}

func (f *fileUploaderService) NewConnection() *session.Session {
	sess := session.New(
		&aws.Config{
			Endpoint: &f.host,
			Region:   aws.String(f.region),
			Credentials: credentials.NewStaticCredentials(
				f.keyID,
				f.keySecret,
				"",
			),
		})
	return sess
}

func (f *fileUploaderService) UploadFile(file []byte) (string, error) {
	sess := f.NewConnection()

	fileName := uuid.New()
	uploader := s3manager.NewUploader(sess)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(f.bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName.String() + ".wav"),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return "", err
	}
	return up.Location, nil
}

func (f *fileUploaderService) DeleteFile(filePath string) error {
	sess := f.NewConnection()

	svc := s3.New(sess)

	strs := strings.Split(filePath, "/")
	fileName := strs[len(strs)-1]

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(f.bucketName), Key: aws.String(fileName)})
	if err != nil {
		return err
	}
	return nil
}
