package buckets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket(client *s3.S3) (string, error) {
	res, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String("cr7bucket1699"),
	})
	if err != nil {
		return "", err
	}

	err = client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String("cr7bucket1699"),
	})

	if err != nil {
		return "", err
	}
	return *res.Location, nil
}
