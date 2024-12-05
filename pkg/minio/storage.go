package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"goresizer.com/m/internal/config"
)

var cfg = config.GetConfig().Minio

var fileID string

func SetFileID(id string) {
	fileID = id
}

func UploadImgFile(img_name string, file multipart.File, size_file int64, contentType string) error {

	minioClient := CreateConncet()

	content, err := minioClient.PutObject(context.Background(), cfg.Storage, img_name, file, size_file,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Uploaded", content.Key, "to", content.Bucket, content.ETag, content.VersionID, content.Size)
	fileID = img_name
	return err
}
func GetImgFile(img_name string) (*minio.Object, error) {
	minioClient := CreateConncet()

	object, err := minioClient.GetObject(context.Background(), cfg.Storage, img_name, minio.GetObjectOptions{})
	if err != nil {
		log.Println("Failed to get object:", err)
		return nil, err
	}

	return object, nil
}

func DownloadImgFile() error {
	minioClient := CreateConncet()

	savePath := fmt.Sprintf("/tmp/download/pp/%s", fileID)

	err := minioClient.FGetObject(
		context.Background(),
		cfg.Storage,
		fileID,
		savePath,
		minio.GetObjectOptions{},
	)

	if err != nil {
		log.Printf("Failed to download file %s: %v", fileID, err)
		return err
	}

	fmt.Printf("File %s downloaded to %s\n", fileID, savePath)
	return nil
}
