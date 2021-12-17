package awsclient

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Api struct {
	session   *session.Session
	s3Service *s3.S3
}

func NewS3API(s *session.Session) *S3Api {
	return &S3Api{
		session:   s,
		s3Service: s3.New(s),
	}
}

func (s3Api *S3Api) GetObjectAsString(bucket, path string) (string, error) {
	output, err := s3Api.GetObject(bucket, path)

	if err != nil {
		return "", err
	}

	return string(output), err
}

func (s3Api *S3Api) GetObject(bucket, path string) ([]byte, error) {

	input := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	var output *s3.GetObjectOutput
	var err error

	output, err = s3Api.s3Service.GetObject(&input)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(output.Body)
}
