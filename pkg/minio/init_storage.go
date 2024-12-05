package storage

import (
	"context"
	"time"

	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"goresizer.com/m/internal/config"
)

func CreateConncet() *minio.Client {
	cfg := config.GetConfig().Minio
	useSSL := false

	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exists, err := minioClient.BucketExists(ctx, cfg.Storage)
	if err != nil {
		log.Fatalf("Error while checking bucket: %v", err)
	}

	if !exists {
		log.Printf("Bucket %s does not exist. Creating...", cfg.Storage)
		err = minioClient.MakeBucket(ctx, cfg.Storage, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Error while creating bucket %s: %v", cfg.Storage, err)
		}
		log.Printf("Bucket %s successfuly created", cfg.Storage)
	}

	return minioClient
}

func CreateBucket() {
	minioClient := CreateConncet()
	cfg := config.GetConfig().Minio

	bucketName := cfg.Storage
	err := minioClient.MakeBucket(context.Background(), cfg.Storage, minio.MakeBucketOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	bucket_are_exist, err := minioClient.BucketExists(context.Background(), cfg.Storage)
	if bucket_are_exist {
		fmt.Println("Bucket: ", bucketName, " created")
	} else {
		fmt.Println("Bucket does not exist.")
	}

}
