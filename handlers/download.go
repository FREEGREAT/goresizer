package handlers

import (
	"fmt"
	"net/http"

	storage "goresizer.com/m/minio"
)

func DownloadImgHandler(w http.ResponseWriter, r *http.Request) {
	// fileID := r.URL.Query().Get("id")
	// if fileID == "" {
	// 	http.Error(w, "Параметр id обов'язковий", http.StatusBadRequest)
	// 	return
	// }

	storage.DownloadImgFile()
	fmt.Fprintf(w, "Файл з id %s завантажено!")
}
