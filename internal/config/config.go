package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"goresizer.com/m/pkg/logging"
)

type Config struct {
	MongoDB struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		AuthDB     string `yaml:"auth_db"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
	Minio struct {
		Endpoint  string `yaml:"endpoint"`
		Storage   string `yaml:"storage"`
		SecretKey string `yaml:"secret_k"`
		AccessKey string `yaml:"access_k"`
	} `yaml:"minio"`
	RabbitMQ struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		QueueName string `yaml:"queuename"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
	} `yaml:"rabbitmq"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read app config")
		instance = &Config{}

		possiblePaths := []string{
			"config.yml",
			"../config.yml",
			"./config.yml",
		}

		var err error
		for _, path := range possiblePaths {
			err = cleanenv.ReadConfig(path, instance)
			if err == nil {
				break
			}
		}

		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
