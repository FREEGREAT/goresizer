package handlers

import (
	"fmt"
	"net/http"
	"time"

	storage "goresizer.com/m/minio"
)

func UploadImgHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "It`s too large", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	storage.UploadImgFile(time.Now().GoString()+handler.Filename, file, handler.Size, handler.Header.Get("Content-Type"))
	fmt.Println("File name \n", handler.Filename)
	fmt.Println("Header \n", handler.Header)

}
