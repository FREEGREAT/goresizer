package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

const bName = "pic-storage"

var fileID string

func UploadImgFile(img_name string, file multipart.File, size_file int64, contentType string) error {

	minioClient := CreateConncet()

	content, err := minioClient.PutObject(context.Background(), bName, img_name, file, size_file,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Uploaded", content.Key, "to", content.Bucket, content.ETag, content.VersionID, content.Size)
	fileID = img_name
	return err
}

func DownloadImgFile() {
	minioClient := CreateConncet()

	savePath := "/tmp/download/pp/" + fileID

	err := minioClient.FGetObject(context.Background(), bName, fileID, savePath,
		minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Check folder ", savePath, ":-)")
	}
}
