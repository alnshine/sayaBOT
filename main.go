package main

import (
	"os"

	"github.com/alnshine/sayaBOT/configs"
	"github.com/alnshine/sayaBOT/internal/api"
	"github.com/alnshine/sayaBOT/internal/repository"
	"github.com/alnshine/sayaBOT/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log := logrus.New()

	if err := configs.InitConfig(); err != nil {
		log.Fatalf("error with reading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with loading env files: %s", err.Error())
	}

	token := os.Getenv("TOKEN")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	api.RunTelegramAPI(log, token, service)
}
