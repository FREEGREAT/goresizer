package storage

import (
	"context"

	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateConncet() *minio.Client {

	s_key := "suHfOm63OLmiTs5Kr1VEjGSbUoPhBv5D9zi8setK"
	endpoint := "localhost:9000"
	accessKeyID := "c5y734tSjVR085kWLUtj"
	secretAccessKey := s_key
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// buckets, err := minioClient.ListBuckets(context.Background())
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	return minioClient
}

func CreateBucket() {
	minioClient := CreateConncet()

	bucketName := "pic-storage"
	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	bucket_are_exist, err := minioClient.BucketExists(context.Background(), bucketName)
	if bucket_are_exist {
		fmt.Println("Bucket: ", bucketName, " created")
	} else {
		fmt.Println("Bucket does not exist.")
	}

}
