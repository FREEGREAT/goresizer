package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	producer "goresizer.com/m/internal/service"
	storage "goresizer.com/m/internal/storage/minio"
)

func UploadImgHandler(w http.ResponseWriter, r *http.Request) {
	resizePercent := r.URL.Query().Get("resizepercent")
	if resizePercent == "" {
		http.Error(w, "Необхідно вказати відсоток зміни розміру", http.StatusBadRequest)
		return
	}

	resizeValue, err := strconv.ParseFloat(resizePercent, 64)
	if err != nil || resizeValue <= 0.0 || resizeValue > 1.0 {
		http.Error(w, "Відсоток зміни розміру має бути числом від 0.1 до 1", http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Файл занадто великий. Максимальний розмір: 10MB", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Помилка при отриманні файлу", http.StatusBadRequest)
		return
	}
	defer file.Close()

	contentType := handler.Header.Get("Content-Type")
	if !isValidImageType(contentType) {
		http.Error(w, "Непідтримуваний тип файлу. Дозволені формати: JPEG, PNG, GIF", http.StatusBadRequest)
		return
	}

	filename := generateUniqueFileName(handler.Filename)

	err = storage.UploadImgFile(filename, file, handler.Size, contentType)
	if err != nil {
		http.Error(w, "Помилка при завантаженні файлу", http.StatusInternalServerError)
		return
	}

	err = producer.PublishMessage(filename, resizeValue)
	if err != nil {
		http.Error(w, "Помилка при відправці повідомлення в чергу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Файл успішно завантажено",
		"filename": filename,
	})
}

func isValidImageType(contentType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/png",
		"image/gif",
	}
	for _, t := range validTypes {
		if t == contentType {
			return true
		}
	}
	return false
}
func generateUniqueFileName(originalFileName string) string {
	timestamp := time.Now().Format("20060102_150405")
	ext := filepath.Ext(originalFileName)

	cleanFileName := strings.ReplaceAll(originalFileName, " ", "_")
	cleanFileName = regexp.MustCompile(`[^a-zA-Z0-9._-]`).ReplaceAllString(cleanFileName, "")

	return fmt.Sprintf("%s_%s%s", timestamp, cleanFileName, ext)
}
