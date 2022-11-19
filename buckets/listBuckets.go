package buckets

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

func ListAllBuckets(client *s3.S3) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		log.Fatal("cannot list buckets")
	}

	fmt.Println("Buckets list:", res.Buckets)
	fmt.Println("Owner name:", *res.Owner.DisplayName)
}
