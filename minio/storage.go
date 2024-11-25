package storage

import (
	"context"
	"fmt"
	"log"

	"goresizer.com/m/ui"

	"github.com/minio/minio-go/v7"
)

const bName = "pic-storage"

func UploadImgFile(img_path string) {


	oName := ui.CreateFileName() + ".jpg"

	minioClient := CreateConncet()

	content, err := minioClient.FPutObject(context.Background(), bName, oName, img_path,
		minio.PutObjectOptions{ContentType: "image/jpg"})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Uploaded", content.Key, "to", content.Bucket, content.ETag, content.VersionID, content.Size)

}

func DownloadImgFile() {
	minioClient := CreateConncet()

	oName := "2.jpg"
	savePath := "/tmp/download/pp/" + oName

	err := minioClient.FGetObject(context.Background(), bName, oName, savePath,
		minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Check folder ", savePath, ":-)")
	}
}
