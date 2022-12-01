package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"github.com/suryateja1698/s3/buckets"
)

var (
	Client *s3.S3
)

func main() {

	location, err := buckets.CreateBucket(Client)
	if err != nil {
		log.Fatal("err in creating bucket", err)
	}
	log.Println("location of bucket created is:", location)

	buckets.ListAllBuckets(Client)
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("couldn't load env")
	}

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if len(accessKey) == 0 || len(secret) == 0 {
		log.Fatalf("Invalid credentials")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
	})
	if err != nil {
		log.Printf("err in creating new session %s", err)
	}

	// check if credentials can be retreived
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("err is:", err)
	}

	// create a new client for s3 with the session
	Client = s3.New(sess)
}
