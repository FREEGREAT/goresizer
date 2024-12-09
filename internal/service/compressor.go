package service

import (
	"bytes"
	"image"
	"image/jpeg"

	"log"

	"github.com/nfnt/resize"

	storage "goresizer.com/m/internal/storage/minio"
)

type BytesReaderFile struct {
	*bytes.Reader
}

func (b *BytesReaderFile) Close() error {
	return nil
}

func Compress(imgName string, resizePercent float64) error {

	obj, err := storage.GetImgFile(imgName)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(obj)
	if err != nil {
		return err
	}

	imgDecode, _, err := image.Decode(buffer)
	if err != nil {
		return err
	}

	m := resize.Resize(
		uint(float64(imgDecode.Bounds().Dx())*resizePercent),
		0,
		imgDecode,
		resize.Bicubic,
	)

	compressedFileName := "compressed_" + imgName

	var compressedBuffer bytes.Buffer
	err = jpeg.Encode(&compressedBuffer, m, nil)
	if err != nil {
		return err
	}

	compressedReader := &BytesReaderFile{
		Reader: bytes.NewReader(compressedBuffer.Bytes()),
	}

	err = storage.UploadImgFile(
		compressedFileName,
		compressedReader,
		int64(compressedBuffer.Len()),
		"image/jpeg",
	)
	if err != nil {
		log.Printf("Failed to upload compressed image: %v", err)
		return err
	}

	return nil
}
