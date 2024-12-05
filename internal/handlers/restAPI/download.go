package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	storage "goresizer.com/m/pkg/minio"
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

	homeDir, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, "Failed to determine home directory", http.StatusInternalServerError)
		return
	}

	downloadDir := filepath.Join(homeDir, "download")
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		err := os.MkdirAll(downloadDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Failed to create download directory", http.StatusInternalServerError)
			return
		}
	}

	localFilePath := filepath.Join(downloadDir, filename)
	fmt.Println("Serving file from:", localFilePath)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, localFilePath)
}
