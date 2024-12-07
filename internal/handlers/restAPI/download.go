package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	storage "goresizer.com/m/internal/storage/minio"
)

func DownloadImgHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	filename = filepath.Base(filename)
	storage.SetFileID(filename)

	err := storage.DownloadImgFile()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to download file: %v", err), http.StatusInternalServerError)
		return
	}

	localFilePath := filepath.Join("/tmp/download/pp", filename)
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		http.Error(w, "File not found after download", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, localFilePath)
}
