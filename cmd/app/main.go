package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"goresizer.com/m/internal/config"
	"goresizer.com/m/internal/handlers/amqp/consumer"
	middleware "goresizer.com/m/internal/handlers/middleware"
	handlers "goresizer.com/m/internal/handlers/restAPI"
	user "goresizer.com/m/internal/storage/db"
	db "goresizer.com/m/internal/storage/mongodb"
	"goresizer.com/m/pkg/logging"
	"goresizer.com/m/pkg/mongodb"
)

const defaultPath = ``
const bName = "pic-storage"

func main() {
	logger := logging.GetLogger()
	cfg := initConfig(&logger)

	fmt.Print(os.Executable())
	mongoClient, err := mongodb.NewClient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		log.Fatal(err)
	}

	storage := db.NewStorage(mongoClient, cfg.MongoDB.Collection, &logger)
	go consumer.Consumer()

	r := initRouter(storage)
	logger.Info("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initConfig(logger *logging.Logger) *config.Config {
	cfg := config.GetConfig()
	if cfg == nil {
		logger.Fatal("Конфігурація не завантажена")
	}
	return cfg
}

func initRouter(storage user.Storage) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handlers.SignUpHandler(storage)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(storage)).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(storage))
	protected.HandleFunc("/upload", handlers.UploadImgHandler).Methods("POST")
	protected.HandleFunc("/download", handlers.DownloadImgHandler).Methods("GET")

	return r
}
