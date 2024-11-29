package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"goresizer.com/m/handlers"
)

const defaultPath = ``
const bName = "pic-storage"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.UploadImgHandler).Methods("POST")
	r.HandleFunc("/download", handlers.DownloadImgHandler).Methods("GET")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
