package bucket

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/config"
)

type Bucket struct {
	name      string
	accessUrl string
	client    *s3.S3
}

func (b *Bucket) Name() string {
	return b.name
}

func (b *Bucket) AccessUrl() string {
	return b.accessUrl
}

// TODO content-type
func (b *Bucket) Upload(filename string, contentType string, fileBytes []byte) error {
	if _, err := b.client.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(fileBytes),
		Bucket:      aws.String(b.name),
		Key:         aws.String(filename),
		ContentType: aws.String(contentType),
	}); err != nil {
		logrus.Error(err)
		return PutFileToBucketFailedError
	}

	return nil
}

func NewBucket(name string, s3Config config.S3Config) (*Bucket, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Config.Id, s3Config.Secret, ""),
		Endpoint:         aws.String(s3Config.Endpoint),
		Region:           aws.String(s3Config.Region),
		DisableSSL:       aws.Bool(s3Config.DisableSSL),
		S3ForcePathStyle: aws.Bool(s3Config.S3ForcePathStyle),
	})

	if err != nil {
		return nil, err
	}

	client := s3.New(sess)
	return &Bucket{name: name, client: client, accessUrl: s3Config.AccessEndpoint + "/" + name}, nil
}
