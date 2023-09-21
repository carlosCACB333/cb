package utils

import (
	"cb/libs"
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func PutObject(file *multipart.FileHeader, bucket string, key string) error {

	// Get S3 Client
	s3Client := libs.GetS3Client()
	// Open file
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// Upload file

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
		Metadata: map[string]string{
			"Content-Type": file.Header.Get("Content-Type"),
		},
	})
	if err != nil {
		return err
	}

	return nil

}

func DeleteObject(bucket string, key string) error {

	// Get S3 Client
	s3Client := libs.GetS3Client()

	// Delete file
	_, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil

}

func GetObject(bucket string, key string) (*s3.GetObjectOutput, error) {

	// Get S3 Client
	s3Client := libs.GetS3Client()

	// Get file
	res, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return res, nil

}
