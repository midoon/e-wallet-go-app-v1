package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
	Redis    Redis
	RabbitMQ RabbitMQ
}

func GetConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error when load env :  %s", err.Error())
	}

	return &Config{
		Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		JWT{
			Key:    os.Getenv("JWT_KEY"),
			Issuer: os.Getenv("JWT_ISSUER"),
		},
		Redis{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		RabbitMQ{
			Username: os.Getenv("MQ_USERNAME"),
			Password: os.Getenv("MQ_PASSWORD"),
			Host:     os.Getenv("MQ_HOST"),
			Port:     os.Getenv("MQ_PORT"),
			User:     os.Getenv("MQ_USER"),
			Exchange: os.Getenv("MQ_EXCHANGE"),
			RKey:     os.Getenv("MQ_RKEY"),
			Queue:    os.Getenv("MQ_QUEUE"),
		},
	}
}
