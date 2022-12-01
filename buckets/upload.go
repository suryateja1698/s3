package buckets

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(s *session.Session) error {
	uploader := s3manager.NewUploader(s)
	fileName := "wall.jpg"

	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", fileName, err)
	}

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("cr7bucket1699"),
		Key:    aws.String("wallpaper"),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed in upload %s , with error %v", fileName, err)
	}

	fmt.Println("uploaded:", output)
	return nil
}

func GetFile(s *session.Session) error {
	filename := "my.jpg"
	downloader := s3manager.NewDownloader(s)

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("cr7bucket1699"),
		Key:    aws.String("wallpaper"),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}

	return nil
}
