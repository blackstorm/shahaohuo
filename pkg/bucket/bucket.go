package bucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/config"
)

type Bucket struct {
	name   string
	client *s3.S3
}

func (b *Bucket) Name() string {
	return b.name
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

func NewBucket(name string, s3Config config.S3Config, createBucketIfNotExist bool) (*Bucket, error) {
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
	input := &s3.HeadBucketInput{
		Bucket: aws.String(name), // 必须
	}

	// TODO 重复代码
	if _, err := client.HeadBucket(input); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				if createBucketIfNotExist {
					if e := createBucket(name, client); e != nil {
						logrus.Error(e)
						return nil, CreateBucketFailedError
					} else {
						logrus.Info("create bucket {} success", name)
					}
				} else {
					return nil, BucketNotExistError
				}
			case ErrCodeNoFound:
				if createBucketIfNotExist {
					if e := createBucket(name, client); e != nil {
						logrus.Error(e)
						return nil, CreateBucketFailedError
					} else {
						logrus.Info("create bucket {} success", name)
					}
				} else {
					return nil, BucketNotExistError
				}
			default:
				logrus.Error(aerr.Error())
			}
		}
	}

	return &Bucket{name: name, client: client}, nil
}

func createBucket(name string, client *s3.S3) error {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(name), // 必须
	}

	if _, e := client.CreateBucket(input); e == nil {
		return readOnlyAnonUserPolicy(name, client)
	} else {
		return e
	}
}

func readOnlyAnonUserPolicy(name string, client *s3.S3) error {
	readOnlyAnonUserPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "AddPerm",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", name),
				},
			},
		},
	}

	policy, _ := json.Marshal(readOnlyAnonUserPolicy)
	_, err := client.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(name),
		Policy: aws.String(string(policy)),
	})

	return err
}
