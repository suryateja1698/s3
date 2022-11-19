package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/suryateja1698/s3/buckets"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String("ap-northeast-1")},
	}))

	// check if credentials can be retreived
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("err is:", err)
	}

	// create a new client for s3 with the session
	client := s3.New(sess)

	location, err := buckets.CreateBucket(client)
	if err != nil {
		log.Fatal("err in creating bucket", err)
	}
	fmt.Println("location of bucket created is:", location)
	buckets.ListAllBuckets(client)
}
